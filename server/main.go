package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsenart/nap"
	io "io/ioutil"
	"jsonrpc"
	"jsonrpc/server/service/user"
	"github.com/go-redis/redis"
	"log"
	"time"
	"runtime"
	//"strings"
	//"net"
	"jsonrpc/registry"
	"os"
	//"strconv"
	"strconv"
)

var db *nap.DB
var client *redis.Client
var serviceAddress string
var weight int

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

type DbOption struct {
	Driver   string
	Host     string
	Port     string
	Dbname   string
	User     string
	Password string
	Charset  string
	Tag      string
}

type DbOptions struct {
	Master DbOption
	Slave  []DbOption
}

type RedisOptions struct {
	Tag      string
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

type RedisDbOptions struct {
	Master RedisOptions
	Slave  []RedisOptions
}

func (self *JsonStruct) Load(filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		return
	}
	datajson := []byte(data)

	err = json.Unmarshal(datajson, v)
	if err != nil {
		return
	}
}

func init() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Usage: ", os.Args[0], "server:port [weight]")
		log.Fatal(os.Args)
	}

	serviceAddress = os.Args[1]
	weight = 1
	var err error

	if len(os.Args) == 3 {
		weight, err = strconv.Atoi(os.Args[2])
		checkError(err)
	}

	kvs := DbOptions{}

	JsonParse := NewJsonStruct()
	JsonParse.Load("config/db.json", &kvs)

	ip := kvs.Master.Host
	port := kvs.Master.Port
	user := kvs.Master.User
	password := kvs.Master.Password
	dbname := kvs.Master.Dbname
	charset := kvs.Master.Charset

	dsns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s;", user, password, ip, port, dbname, charset)

	for _, v := range kvs.Slave {
		ip = v.Host
		port = v.Port
		user = v.User
		password = v.Password
		dbname = v.Dbname
		charset = v.Charset
		dsns += fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s;", user, password, ip, port, dbname, charset)
	}
	dsns = string([]rune(dsns)[:len(dsns) - 1])

	db, err = nap.Open("mysql", dsns)

	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println("%s error ping database: %s", err.Error())
		return
	}

	db.SetMaxIdleConns(100)
	tickDbPing()

	client = createClient()

	registry.Inject(client)

	registry.Register(serviceAddress, weight)
	return
}

func createClient() *redis.Client {
	kvs := RedisDbOptions{}
	var err error

	JsonParse := NewJsonStruct()
	JsonParse.Load("config/redis.json", &kvs)

	host := kvs.Master.Host
	port := kvs.Master.Port
	password := kvs.Master.Password
	poolSize := kvs.Master.PoolSize
	dbName := kvs.Master.DB

	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       dbName,
		PoolSize: poolSize,
	})

	_, err = client.Ping().Result()
	//fmt.Println(pong, err)
	checkError(err)

	return client
}

func tickDbPing() {
	ticker := time.NewTicker(time.Second * 8)
	go func() {
		for range ticker.C {
			db.Ping()
		}
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU());
	user.Register(db)
	user.InjectSA(serviceAddress)
	fmt.Println("rpc server running.")
	jsonrpc.ListenAndServe("tcp", serviceAddress)
}
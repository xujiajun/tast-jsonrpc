package main

import (
	"jsonrpc"
	registryService "jsonrpc/registry"
	"fmt"
	"github.com/go-redis/redis"
	io "io/ioutil"
	"encoding/json"
)

var client *redis.Client

type Registry struct {
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

func (registry *Registry) GetIp(_, reply *string) error {
	*reply = registryService.GetIp()
	return nil
}

type JsonStruct struct {
}

func NewJsonStruct2() *JsonStruct {
	return &JsonStruct{}
}

func (self *JsonStruct) Load2(filename string, v interface{}) {
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
	kvs := RedisDbOptions{}

	JsonParse := NewJsonStruct2()
	JsonParse.Load2("config/redis.json", &kvs)

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

	_, err := client.Ping().Result()
	//fmt.Println(pong, err)
	if err != nil {
		panic(err)
	}
	registryService.Inject(client)
	registryService.HealthCheck()

}
func main() {
	jsonrpc.Register(new(Registry))
	fmt.Println("registry server running.")
	jsonrpc.ListenAndServe("tcp", ":1231")
}
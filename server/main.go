package main

import (
	"database/sql"
	json "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	io "io/ioutil"
	"jsonrpc"
	"jsonrpc/server/service/user"
	"log"
	"time"
)

var db *sql.DB

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
}

type DbOptions struct {
	Master DbOption
	Slave  DbOption
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
	v := DbOptions{}
	var err error

	JsonParse := NewJsonStruct()
	JsonParse.Load("config/db.json", &v)

	drive := v.Master.Driver
	ip := v.Master.Host
	port := v.Master.Port
	user := v.Master.User
	password := v.Master.Password
	dbname := v.Master.Dbname
	charset := v.Master.Charset

	db, err = sql.Open(drive, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, ip, port, dbname, charset))
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println("%s error ping database: %s", err.Error())
		return
	}
	db.SetMaxIdleConns(100)
	tickDbPing()
	return
}

func tickDbPing() {
	ticker := time.NewTicker(time.Second * 8)
	go func() {
		for range ticker.C {
			db.Ping()
		}
	}()
}

func main() {
	user.Register(db)
	jsonrpc.ListenAndServe("tcp", ":1234")
}

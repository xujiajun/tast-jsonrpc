package main

import (
	"database/sql"
	json "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	io "io/ioutil"
	"jsonrpc"
	"jsonrpc/server/service/user"
	"log"
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

	db, err = sql.Open(drive, user+":"+password+"@tcp("+ip+":"+port+")/"+dbname+"?charset="+charset)
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	db.SetMaxIdleConns(100)
	return
}

func main() {
	user.Register(db)
	jsonrpc.ListenAndServe("tcp", ":1234")
}

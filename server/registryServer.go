package main

import (
	"jsonrpc"
	registryService "jsonrpc/registry"
	"jsonrpc/common"
	"fmt"
	"github.com/go-redis/redis"
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

func init() {
	kvs := RedisDbOptions{}

	JsonParse := common.NewJsonStruct()
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
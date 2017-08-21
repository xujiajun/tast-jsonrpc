package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"jsonrpc"
	"jsonrpc/common"
	registryService "jsonrpc/registry"
)

var client *redis.Client

// Registry definition
type Registry struct {
}

// RedisOptions definition
type RedisOptions struct {
	Tag      string
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

// RedisDbOptions definition
type RedisDbOptions struct {
	Master RedisOptions
	Slave  []RedisOptions
}

// GetIP publishes the receiver's methods GetIP in the DefaultServer.
func (registry *Registry) GetIP(_, reply *string) error {
	*reply = registryService.GetIP()
	return nil
}

func init() {
	kvs := RedisDbOptions{}

	JSONParse := common.NewJSONStruct()
	JSONParse.Load("config/redis.json", &kvs)

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

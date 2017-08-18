package main

import (
	"jsonrpc"
	registryService "jsonrpc/registry"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

type Registry struct {
}

func (registry *Registry) GetIp(_, reply *string) error {
	*reply = registryService.GetIp()
	return nil
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 40,
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
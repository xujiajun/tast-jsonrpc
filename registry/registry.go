package registry

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/smallnest/weighted"
	"net"
	"strconv"
	"time"
)

//Registry definition interface
type Registry interface {
	Register()
	RegisterService()
	UnRegister()
	GetIP()
	HealthCheck()
}

var client *redis.Client

//Register to implement method Register in interface Registry
func Register(ip string, weight int) {
	client.HSet("ips", ip, weight)
	client.HSet("ips_all", ip, weight)
	client.HSet("s_status", ip, 1)
}

//RegisterService to implement method RegisterService in interface Registry
func RegisterService(ip string, serviceName string) {
	client.HSet("services", ip, serviceName)
}

//UnRegister to implement method UnRegister in interface Registry
func UnRegister(ip string) {
	client.HDel("ips", ip)
	client.HSet("s_status", ip, -1)
}

//Inject redis client for invoking
func Inject(redisClient *redis.Client) {
	client = redisClient
}

//GetIP to implement method GetIP in interface Registry
func GetIP() string {
	var w = weighted.W1{}
	hash, err := client.HGetAll("ips").Result()
	if err != nil {
		panic(err)
	}

	for ip, weight := range hash {
		weight, err := strconv.Atoi(weight)
		if err != nil {
			panic(err)
		}
		w.Add(ip, weight)
	}

	return fmt.Sprintf("%s", w.Next())
}

//HealthCheck to implement method HealthCheck in interface Registry
func HealthCheck() {
	ticker := time.NewTicker(time.Second * 1)
	hash, err := client.HGetAll("ips").Result()
	if err != nil {
		panic(err)
	}
	go func() {
		for range ticker.C {
			for ip := range hash {
				tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
				_, err = net.DialTCP("tcp", nil, tcpAddr)
				if err != nil {
					//panic("xxx")
					UnRegister(ip)
				}
			}
		}
	}()
}

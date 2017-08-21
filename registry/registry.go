package registry

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/smallnest/weighted"
	"net"
	"strconv"
	"time"
)

type Registry interface {
	Register()
	RegisterService()
	UnRegister()
	GetIp()
	HealthCheck()
}

var client *redis.Client

func Register(ip string, weight int) {
	client.HSet("ips", ip, weight)
	client.HSet("ips_all", ip, weight)
	client.HSet("s_status", ip, 1)
}

func RegisterService(ip string, serviceName string) {
	client.HSet("services", ip, serviceName)
}

func UnRegister(ip string) {
	client.HDel("ips", ip)
	client.HSet("s_status", ip, -1)
}

func Inject(redisClient *redis.Client) {
	client = redisClient
}

func GetIp() string {
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

func HealthCheck() {
	ticker := time.NewTicker(time.Second * 1)
	hash, err := client.HGetAll("ips").Result()
	if err != nil {
		panic(err)
	}
	go func() {
		for range ticker.C {
			for ip, _ := range hash {
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

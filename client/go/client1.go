package main

import (
	//"fmt"
	"jsonrpc"
	//"log"
	//"time"
	"fmt"
	"log"
	"time"
)

func main() {

	var reply, x string

	client, err := jsonrpc.NewClient("tcp", "127.0.0.1:1231")
	err = client.Call("Registry.GetIP", x, &reply)
	fmt.Printf("Registry.GetIp: %s\n", reply)
	client, err = jsonrpc.NewClient("tcp", reply)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	start := time.Now().UnixNano()
	//for n := 0; n <= 10000; n++ {
	err = client.Call("User.GetUser", x, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("User.getUser: %s\n", reply)
	//}
	//time.Sleep(time.Second * 2)
	end := time.Now().UnixNano()
	sub := (end - start) / 1000000000
	fmt.Printf("执行时间: %d（s）\n", sub)
	//fmt.Printf("QPS: %d \n", 10000 / sub);
}

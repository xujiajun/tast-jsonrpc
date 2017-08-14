package main

import (
	"fmt"
	"jsonrpc"
	"log"
)

func main() {
	client, err := jsonrpc.NewClient("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply, x string
	err = client.Call("User.GetUser", x, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("User.getUser: %s\n", reply)
}

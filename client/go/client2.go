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

	var reply int
	mapD := map[string]float64{"num1": 5, "num2": 7}
	err = client.Call("MyMath.Add", mapD, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("MyMath.Add: %d\n", reply)
}

package main

import (
	"jsonrpc"
)

type MyMath struct {
}

func (mm *MyMath) Add(num map[string]float64, reply *float64) error {
	*reply = num["num1"] + num["num2"]
	return nil
}

func (mm *MyMath) Sub(num map[string]string, reply *string) error {
	*reply = num["num1"] + num["num2"]
	return nil
}

func main() {
	jsonrpc.Register(new(MyMath))
	jsonrpc.ListenAndServe("tcp", ":1234")
}

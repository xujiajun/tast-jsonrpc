package jsonrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ListenAndServe(host, port string) {
	listener, err := net.Listen(host, port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func Register(server interface{}) {
	rpc.Register(server)
}

func NewClient(network, address string) (*rpc.Client, error) {
	client, err := jsonrpc.Dial(network, address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client, err
}

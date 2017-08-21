package jsonrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//ListenAndServe  Listen announces on the local network address and process requests
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

//Register publishes the receiver's methods in the DefaultServer.
func Register(server interface{}) {
	rpc.Register(server)
}

//NewClient connects to a JSON-RPC server at the specified network address.
func NewClient(network, address string) (*rpc.Client, error) {
	client, err := jsonrpc.Dial(network, address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client, err
}

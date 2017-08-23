# Tast-JSONRPC [![Go Report Card](https://goreportcard.com/badge/github.com/xujiajun/tast-jsonrpc)](https://goreportcard.com/report/github.com/xujiajun/tast-jsonrpc)

The Go language implementation of jsonrpc framework. Support PHP、Golang、Java client .



## Architecture

Tast-JSONRPC contains four roles : RPC Server, RPC Client, Registry and Admin console.

* RPC Server registers services on the Registry

* RPC Client select a server from the server list returned from the Registry.

* Registry manage server node list（register、remove、health check、Load balancing etc.）

* Admin console as a role of service governance（run、stop server、modify weight etc.）

<img src ="https://raw.githubusercontent.com/xujiajun/tast-jsonrpc/master/docs/rpc-architecture-en.png">


## Features

* Based on net/rpc 、net/rpc/jsonrpc
* Cross-language client support (Golang、Java、PHP)
* Data layer（now support mysql master-slave architecture）
* Service dicovery (based on redis)
* Server Healthcheck
* Load banlancer (now support RoundRobin)
* Scalable
* Service governance
* ...


## Server

visit [this](https://github.com/xujiajun/tast-jsonrpc/blob/master/server/main.go)

## Client

* [Go client](https://github.com/xujiajun/tast-jsonrpc/blob/master/client/go/client1.go)

* [Java client](https://github.com/xujiajun/tast-jsonrpc/blob/master/client/java/Client.java)

* [PHP client](https://github.com/xujiajun/tast-jsonrpc/blob/master/client/php/client.php)

## Registry Server

visit [this](https://github.com/xujiajun/tast-jsonrpc/blob/master/server/registryServer.go)

## Admin Console

Link: https://github.com/xujiajun/jsonrpc-admin

### Screenshot

<img src="https://raw.githubusercontent.com/xujiajun/jsonrpc-admin/master/docs/rpc-console.png">

## License
The Tast-JSONRPC is open-sourced software licensed under the [MIT Licensed](http://www.opensource.org/licenses/MIT)

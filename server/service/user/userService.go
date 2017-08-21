package user

import (
	"github.com/tsenart/nap"
	"jsonrpc"
	"jsonrpc/registry"
	"jsonrpc/server/service/user/dao"
)

var db *nap.DB
var serviceAddress string

// ServiceName definition
const ServiceName = "User"

// User definition
type User struct {
}

// UserService definition
type UserService interface {
	GetUser()
}

//GetUser get all user info
func (u *User) GetUser(x, reply *string) error {
	*reply = dao.GetUser(db)
	return nil
}

//Register publishes the receiver's methods User in the DefaultServer.
func Register(DB *nap.DB) {
	jsonrpc.Register(new(User))
	registry.RegisterService(serviceAddress, ServiceName)
	db = DB
}

//InjectSA inject server address into registry
func InjectSA(ServiceAddress string) {
	serviceAddress = ServiceAddress
}

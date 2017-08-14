package user

import (
	"jsonrpc/server/service/user/dao"
	"jsonrpc"
	"database/sql"
)

var db *sql.DB;

type User struct {
}

func (u *User) GetUser(x,reply *string) error {
	*reply = dao.GetUser(db)
	return nil
}

func Register(DB *sql.DB)  {
	jsonrpc.Register(new(User))
	db = DB
}
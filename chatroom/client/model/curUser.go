package model
import (
	"go_code/chatroom/common/model"
	"net"
)

type CurUser struct{
	model.User
	Conn net.Conn 
}
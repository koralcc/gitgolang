package model

import (
	"go_code/project1/94chatroom/common/message"
	"net"
)

//因为在客户端，很多地方要
type CurUser struct {
	Conn net.Conn
	message.User
}

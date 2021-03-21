package process

import (
	"fmt"
	"go_code/project1/94chatroom/client/model"
	"go_code/project1/94chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User)
var CurUser model.CurUser //在用户登陆成功后完成对CurUser的初始化

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

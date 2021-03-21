package process

import (
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/common/message"
	"go_code/project1/94chatroom/server/model"
	"go_code/project1/94chatroom/server/utils"
	"net"
)

type UseProcess struct {
	Conn net.Conn
	//增加一个字段表示该Conn是哪个用户
	UserId int
}

//处理登陆请求
func (this *UseProcess) ServerProcessLogin(mes *message.Message) (err error) {
	fmt.Println("客户端传入数据：", mes)
	//1.反序列化数据
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshao fail err=", err)
		return
	}
	//2.声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//3.执行数据逻辑
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	fmt.Println(user, "登陆成功了.")
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

		//这里先测试成功，然后再根据返回具体错误
	} else {
		loginResMes.Code = 200
		//这里登陆成功后，将信息放到userMgr中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他的在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前用户的id，放入到loginResMes.UserIds
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登陆成功")
	}
	//验证用户名和密码
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册再使用"
	// }
	//4.封装返回消息体
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail:", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail:", err)
		return
	}
	//5.发送客户端响应消息
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//处理注册请求
func (this *UseProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//2.声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	//再声明一个LoginResMes
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail:", err)
		return
	}
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//这里我们编写通知所有在线用户的方法
//userId要通知其他的在线用户，我上线
func (this *UseProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UseProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err", err)
		return
	}
}

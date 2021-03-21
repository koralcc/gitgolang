package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/client/utils"
	"go_code/project1/94chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) err:", err)
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err:", err)
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err ：", err)
	}

	mes, err = tf.ReadPkg() //mes 就是RegisterResMesType
	if err != nil {
		fmt.Println("tf.ReadPkg() err:", err)
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//可以有多个main文件，但是只能有一个main方法，都是main包
//多个main文件的编译，需要编译整个文件夹,在gopath下，go build -o client.exe go_code/project1/94chatroom/client
//编译详情参考27
//写一个函数，完成用户登陆
func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	//下一个就是定协议
	// fmt.Printf("userId= %d userPwd=%s", userID, userPwd)
	// return nil
	//1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("jsong.Marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal err =", err)
		return
	}
	//7.到这个时候，data就是我们要发送的消息
	//7.1先把data的长度发送给服务器
	//先获取到data的长度 -> 转成一个表示长度的byte切片，因为write方法是发送一个byte切片
	var pkgLen = uint32(len(data))
	var buf [4]byte                              //2xy32 == 3G
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //把数字放到一个字节数组中
	//发送长度
	n, err := conn.Write(buf[:4])
	//发送四个字节过去
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度=%d，内容=%s", len(data), string(data))
	//发送消息本身
	fmt.Printf("客户端，发送消息的长度=%d 内容=%s:\n", len(data), string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write(data) fail", err)
		return
	}

	//处理服务器的返回信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err:", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("登陆成功")
		//可以显示但当前在线用户列表
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UserIds {
			//不显示自己
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端的onlineUsers并初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//保持连接：这里需要新建一个协程，保持与服务端的连接，持续读取服务器端往回写的数据
		go serverProcessMes(conn)
		//1.显示我们登陆成功的菜单【循环】
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	// else if loginResMes.Code == 500 {
	// 	fmt.Println(loginResMes.Error)
	// }
	return nil
}

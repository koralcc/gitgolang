package main

import (
	"fmt"
	"go_code/project1/94chatroom/common/message"
	"go_code/project1/94chatroom/server/process"
	"go_code/project1/94chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//总控
//路由消息请求
//优化效率：缓存+算法
//优化结构：分层
//因为进来的消息会有很多种类，需要在服 务器连接处理中分层(相当于控制器)
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆逻辑
		up := &process.UseProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册逻辑
		up := &process.UseProcess{
			Conn: this.Conn,
		}
		up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个SmsProcess实例
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理......")
	}
	return
}

func (this *Processor) process() (err error) {
	// defer this.Conn.Close()
	//读客户端发送的信息
	for {
		//1.读取数据包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				//客户端关闭 use of closed network connection
				fmt.Printf("客户端退出[%v]，服务端也退出\n", this.Conn.RemoteAddr())
			} else {
				fmt.Println("readPkg other err:", err)
			}
			return err
		}
		//2.处理数据包
		err = this.serverProcessMes(&mes)
	}
}

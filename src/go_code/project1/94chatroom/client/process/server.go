package process

import (
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/client/utils"
	"go_code/project1/94chatroom/common/message"
	"net"
	"os"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("--------恭喜xxx登陆成功----------")
	fmt.Println("--------1. 显示在线用户列表----------")
	fmt.Println("--------2. 发送消息----------")
	fmt.Println("--------3. 信息列表----------")
	fmt.Println("--------4. 退出系统----------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	//因为，我们总会使用到SmsProcess实例，所以我们定义在外面
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户")
	case 2:
		fmt.Println("你想对大家说点什么:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0) //给服务器一个关闭的反馈
	default:
		fmt.Println("你输入的选项不正确......")
	}
}

//和服务端保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务端发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err:", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //有人群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了位置的消息类型")
		}
		//fmt.Println("mes=", mes)
	}
}

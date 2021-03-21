package process

import (
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/common/message"
	"go_code/project1/94chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的onlineUsers.map[int]*UserProcess
	//取出mes的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//这里，过滤出自己，不要发送
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}

}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err =", err)
	}
}

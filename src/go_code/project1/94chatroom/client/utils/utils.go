package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project1/94chatroom/common/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //这是传输时，使用缓存
}

//读取消息函数
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	//1.读长度头
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read(head) err", err)
		//打印错误由方法调用者
		return
	}
	//2.读消息内容
	//[]byte => uint32
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen]) //因为read是阻塞的，一个消息客户端发两次，服务端read两次即可
	if n != int(pkgLen) || err != nil {
		return
	}

	//！！这里mes一定要传地址
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}

	return
}

//发送消息
func (this *Transfer) WritePkg(data []byte) (err error) {
	//1.先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	//var buf [4]byte
	//int => []byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//2.发送数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

package main

import (
	"fmt"
	"go_code/project1/94chatroom/server/model"
	"net"
	"time"
)

func initUserDao() {
	//这里的poo本身就是一个全局变量
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao() //必须要先初始化连接池，再初始化Dao
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦链接成功，则启动一个协撑和客户端保持通讯
		go process2(conn)
	}
}

//处理和客户端通讯
func process2(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误：", err)
		return
	}
}

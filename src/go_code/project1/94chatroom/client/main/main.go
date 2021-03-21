package main

import (
	"fmt"
	"go_code/project1/94chatroom/client/process"
	"os"
)

//定义两个变量，用户id和密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	//var loop = true
	for {
		fmt.Println("------------------------欢迎登陆多人聊天系统------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")
		//阻塞等待用户输入
		fmt.Scanf("%d\n", &key)
		//fmt.Println(key)
		switch key {
		case 1:
			//说明用户要登陆
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登陆
			//1
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户逻辑")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
	//二级菜单
	//根据用户的输入，显示新的菜单
	// if key == 1 {
	// 	//说明用户要登陆
	// 	fmt.Println("请输入用户的id")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("请输入用户的密码")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	//先把登陆的函数，写到另外一个文件
	// 	//login(userId, userPwd) //一个main包下有多个文件时，运行，需要编译后再执行，否则会报错，找不到这个方法
	// 	// if err != nil {
	// 	// 	fmt.Println("登陆失败")
	// 	// } else {
	// 	// 	fmt.Println("登陆成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑")
	// }
}

//一个main包下有多个文件时，运行，需要编译后再执行，否则会报错，找不到这个方法
// go build -o bin/client.exe .\src\proj\94chatroom\client
// windows设置环境变量
// 当前cmd窗口
// set GOPATH B:\go
// 系统空间：
// setx -m B:\go
// 用户空间：
// setx GOPATH B:\go

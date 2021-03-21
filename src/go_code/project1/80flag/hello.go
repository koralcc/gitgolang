package main

import (
	"flag"
	"fmt"
)

func main() {
	//1.定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//2.注册变量
	//&user，就是接收用户命令行中输入的-u后面的参数值
	//"u" ，就是-u指定参数
	//""，默认值
	//"用户名，默认为空"，说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "用户名，默认为空")
	flag.StringVar(&host, "host", "localhost", "用户名，默认为空")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//3.
	//很重要的操作，转换
	flag.Parse()
	//从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回Errhelp

	fmt.Printf("user=%v\npwd=%v\nhost=%v\nport=%v", user, pwd, host, port)
}

// flag包
// flag包用来解析命令行参数
// 说明：前面的方式是比较原生的方式，对解析参数不是特别的方便特别是带有参数形式的命令行
// 比如：mysql -u koral -p 123 这样形式的命令行，go设计者给我们提供了flag包，可以方便的解析命令行参数，而且参数顺序可以随意

// intVar stringVar

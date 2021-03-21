package main

import "fmt"

func main() {
	// var num1 = 10
	// var num2 = 0
	// fmt.Println(num1 / num2) //panic: runtime error: integer divide by zero
	// //报错后,程序退出,函数直接结束,获取错误的意义:(1)提醒错误给管理员(2)让下面的代码继续执行
	// fmt.Println("helloword")

	//引入异常后
	test()
	fmt.Println("main helloword")
}

func errHandle() {
	err := recover() //recover()内置函数,可以捕获到异常
	if err != nil {
		fmt.Println("error:", err)
	}
}

func test() {

	// defer func() {
	// 	err := recover() //recover()内置函数,可以捕获到异常
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 	}
	// }()
	defer errHandle()
	var num1 = 10
	var num2 = 0
	fmt.Println(num1 / num2) //做了异常捕获后,本方法报错行后的代码不执行,main方法调用test()后的正常执行

	fmt.Println("test helloword") //以方法为单位,做了异常捕获后,报错后面的语句都不执行
}

//Golang错误处理
// 1)Golang追求简洁优雅,所以,Go语言不支持传统的try...catch...finally这种处理
// 2)Golang中引入异常的处理方式:defer+recover
// 3)这几个异常的使用场景可以这样简单描述:Go中可以抛出一个panic的异常,然后在defer中通过recover捕获这个异常,然后正常处理

//错误处理的好处:程序不会轻易挂掉,如果加入预警代码,就可以让程序更加健壮.

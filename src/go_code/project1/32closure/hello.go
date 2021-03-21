package main

import (
	"fmt"
	"strings"
)

func main() {
	f := addUpper()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	//当反复调用f时，因为n只初始化一次，因此每调用一次就进行一次累计
	//我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引用到的变量共同构成闭包
	var picName = "hello.jpg"
	var picName2 = "hello~"
	f2 := makeSuffix(".jpg") //一次传入反复使用
	fmt.Printf("文件名%v的全名是%v", picName, f2(picName))
	fmt.Printf("文件名%v的全名是%v", picName2, f2(picName2))
}

//累加器
func addUpper() func(int) int {

	//*********************闭包体
	//闭包内引用的的变量n,str会保持状态, 运算后会保存运算值给下一次调用使用
	var n = 10
	var str string = "hello "
	return func(x int) int {
		n = n + x
		str += "a"
		fmt.Println(str)
		return n
	}
	//*********************闭包体
}

//练习

//编写一个程序
// 1)编写一个makeSuffix(suffix string)可以接收一个文件后缀名(比如.jpg),并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg),则返回 文件名.jpg,如果已经有.jpg后缀，则返回源文件名
// 3)要求使用闭包方式完成
// 4)string.HasSuffix,判断字符串是否含有某个后缀的方法

//传统的实现方法,每次都要传入相同的suffix参数
// func makeSuffix(name string, suffix string) string {
// 	if strings.HasSuffix(name, suffix) {
// 		return name
// 	}
// 	return name + suffix
// }

func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
	//返回的匿名函数与 suffix参数形成闭包
}

//闭包
//闭包就是一个函数和与其相关的引用环境组合而成的一个整体(实体)
//闭包内引用的的变量(n,str)会保持状态, 运算后会保存运算值给下一次调用使用

//1)返回的函数和makeSuffix(suffix string)的suffix变量组成一个闭包，因为返回的函数引用到suffix这个变量
//2)我们体会一下闭包的好处，如果使用传统方法，也可以轻松实现这个功能，但是传统方法需要每次都传入后缀名，比如.jpg,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用，

//闭包有一个最大的好词，把匿名函数引用到的变量保留下来，进行下一次调用的时候继续使用，而普通函数(不考虑全局变量实现)则会在函数调用完后释放掉栈

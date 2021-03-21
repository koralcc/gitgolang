package main

import "fmt"

//全局匿名函数
var mquan = func(a int, b int) int {
	return a + b
}

func main() {
	var a int = 10
	var b int = 20

	//匿名函数的调用方式
	//第一种调用，直接在后面加括号传参即可
	c := func(a int, b int) int {
		return a + b
	}(a, b) //匿名函数后加括号，即为匿名函数的即定义即调用
	fmt.Println(c)
	//第二种调用，定义方法变量进行调用
	m := func(a int, b int) int {
		return a + b
	}
	fmt.Println(m(a, b))

	//全局匿名函数
	fmt.Println(mquan(a, b))
}

//匿名函数
// Go支持匿名函数，如果某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用
// 1)匿名函数的两种调用方式
//   (1)定义匿名函数时就直接调用
//   (2)将匿名函数赋给一个变量(函数变量),再通过该变量来调用匿名函数
// 2)全局匿名函数
// 如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以再程序有效

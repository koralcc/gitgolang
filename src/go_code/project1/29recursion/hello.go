package main

import "fmt"

func main() {
	test(4)
	fmt.Println(getFblq(3))
	fmt.Println(getFn1(3))
	fmt.Println(getMonkeyTao(1))
}

func test(n int) {
	if n > 2 {
		n-- //递归必须向退出递归条件逼进，否则就会无限循环
		test(n)
	}
	fmt.Println("n-", n)
}

//返回第n个斐波拉契数
func getFblq(n int) int {
	//斐波拉契数 1 1 2 3 5 8 13  f(n) = f(n-1) + f(n-2) ，由此可见n必须大于2
	fmt.Println("#####################斐波拉契数########################")
	if n == 1 || n == 2 {
		return 1
		//fmt.Printf("第%v个数字为%v", which, 1)
	} //else {
	return getFblq(n-1) + getFblq(n-2)
	//}
}

//已知f(1)=3;f(n)=2*f(n-1)+1,求f(n)的值
func getFn1(n int) int {
	fmt.Println("#####################函数fn=2fn-1+1########################")
	if n == 1 {
		return 3 //递归退出条件
	} //else {
	//if block ends with a return statement, so drop this else and outdent its block
	return 2*getFn1(n-1) + 1 //此处n-1即为退出递归递减的操作
	//}
}

//猴子吃桃问题，有一堆桃子，每天吃一堆的一半加一个，到第十天时，发现只有一个桃子了，问题：最初有几个桃子
// f(n) = f(n-1)/2 -1，现在知道的是第十天的数字，要想递归往第十天，则需要反过来得到 fn=多少的fn+1的表达式  f(n) = （f(n+1)+1） *2
// f(10) = 1
//第n天有几个桃子
func getMonkeyTao(n int) int {
	fmt.Println("#####################猴子吃桃问题########################")
	if n == 10 {
		return 1
	} //else {
	return (getMonkeyTao(n+1) + 1) * 2 //往递归结束方向
	//}
}

//函数的递归调用
// 基本介绍：函数体内又调用了函数本身。
// 函数递归需要遵守的重要原则
// 1)执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
// 2)函数的局部变量是独立的，不会相互受影响
// 3)递归必须向**退出递归**的条件逼近(这个是递归最重要的)，否则就是无限递归，死循环
// 4)当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者return时，该函数在内存里面本身也被系统销毁

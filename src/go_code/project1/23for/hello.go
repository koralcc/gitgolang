package main

import "fmt"

func main() {
	//for的三种实现方式
	//1
	for i := 0; i < 10; i++ {
		fmt.Println("hello word ")
	}
	//2
	var i int = 0
	for i < 10 {
		fmt.Println("hello word ~")
		i++
	}
	//3
	var i2 int = 0
	for {
		if i2 < 10 {
			fmt.Println("hello word i2")
		} else {
			break
		}
		i2++
	}

	//对字符串遍历，1
	var s string = "hello 中国"
	a := []rune(s) //对于有中文的字符串，必须要转成rune才可以，否则出现乱码int32
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d=%c \n", i, a[i]) //6=中，7=国
	}
	//对字符串的遍历，2//推荐用此种方法
	for i, v := range s { //for range按照字符的方式遍历
		fmt.Printf("~~%d=%c \n", i, v) //6=中，9=国
	}

}

//for循环
// 1)语法格式
// for 循环变量初始化;循环条件;循环变量迭代{
// 	循环操作(语句)
// }
// 2)如上语法格式说明
// (1)对for循环来说，有四个要素
// (2)循环变量初始化
// (3)循环条件
// (4)循环操作(语句)，也叫循环体
// (5)循环变量迭代
// 3)for循环执行的顺序说明
// (1)执行循环变量初始化，比如 i=1
// (2)执行循环条件，比如 i<=10
// (3)如果循环条件为真，就执行循环操作(循环体)：比如：fmt.Println()
// (4)执行循环变量迭代，比如i++
// (5)反复执行2，3，4步骤，直到循环条件为false，就退出for循环

// //for 循环的注意事项和细节说明
// 1)循环条件是返回一个布尔值的表达式
// 2)for循环的第二种使用方式
// for 循环判断条件{
// 	//循环执行语句    //对于for循环的四要素中只有循环判断条件是必须要的，不写默认为true
// }
// 3)for循环的第三种使用方式
// for{
// 	//循环执行语句
// }
// 上面的写法等价于for;;{} 是一个无限循环，通常需要配合break语句使用
// 4)Golang 提供for-range的方式，可以方便遍历字符串和数组，对字符串的遍历，建议用for-range

// for循环的while和do while的实现
// 1)while实现(先判断)
// for{
// 	if 循环条件表达式{
// 		break//跳出for循环
// 	}
// 	循环操作
// 	循环变量迭代
// }
// 2)do while实现(先执行，故do while至少执行一次)
// for{
// 	循环操作
// 	循环变量迭代
// 	if 循环条件表达式{
// 		break//跳出for循环
// 	}
// }

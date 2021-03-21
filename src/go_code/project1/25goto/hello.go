package main

import "fmt"

func main() {

	fmt.Println("f1")
	fmt.Println("f2")
	fmt.Println("f3")
	goto f6
	fmt.Println("f4")
	fmt.Println("f5")
f6:
	fmt.Println("f6")
	// loop:
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 		if i == 2 {
	// 			goto loop // 这种直接goto到for循环 相当于重新开始循环 i :=0,而不像cotinue一样会从迭代循环变量i++开始
	// 		}
	// 	}

}

//跳转控制语句goto
// 1)Go语言的goto语句可以无条件地转移到程序中指定的行
// 2)goto语句通常与条件语句配合使用，可用来实现条件的转移跳出循环体等功能
// 3)在Go程序设计中一般不主张使用goto语句，以免照成程序流程的混乱，使理解和调试程序都困难

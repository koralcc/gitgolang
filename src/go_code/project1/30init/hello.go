package main

import (
	"fmt"
	"go_code/project1/30init/utils"
)

var a int = test()

func main() {
	fmt.Println("***main")
	fmt.Println("***utils.Name", utils.Name)
	fmt.Println("***utils.Age", utils.Age)
}

func init() {
	fmt.Println("****main package init")
}

func test() int {
	fmt.Println("***test")
	return 1
}

//init函数
//每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数前调用
//细节
//  1)如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是变量定义 -> init函数 -> main函数
//  2)init函数最主要的作用，就是完成一些初始化的工作

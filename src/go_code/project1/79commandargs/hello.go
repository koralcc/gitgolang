package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的所有参数有", len(os.Args))
	//遍历os.Args切片，就可'以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("arg[%v]=%v\n", i, v)
	}
}

//命令行参数读取
//我们希望能够获取到命令行输入的各种参数，该如何处理
//os.Args是一个string的切片，用来存储所有的命令行参数(这种方法有很强的局限性，必须按照顺序输入)

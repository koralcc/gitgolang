package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1)创建一个新文件，写入内容 5句 hello word
	//1.打开文件 D:/abc.txt
	filePath := "D:/abc.txt"
	//2.使用写和创建的模式打开
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) //0666只在linux/unix系统中有效
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	//准备写入的5句话
	str := "hello word\n"
	//3.写入时，使用带有缓存的*writer(输出流)
	wr := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		wr.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriteString方法时，其实是先写入
	//到缓存的，所以需要调用Flush方法，将缓存的数据正真写入到文件中，否则文件中会没有数据
	//4.
	wr.Flush()
}

//写文件
// 1)创建一个文件，并将hello word写入文件中
// 方法：os.OpenFile(),bufio.NewWriter(),writer.WriteString(),write.Fush()
// func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
// os.OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
// bufio.NewWriter(),缓冲区输出流，适用于大文件写入

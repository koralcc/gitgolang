package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//概念说明：file的叫法
	//1.file叫file对象
	//2.file叫file指针
	//3.file叫file文件句柄
	file, err := os.Open("D:/test.txt") //以读的模式打开文件
	if err != nil {
		fmt.Println("打开文件错误", err) //如果没有这个文件The system cannot find the file specified.
	}
	fmt.Println("打开的文件", file)
	//函数退出时，要及时的关闭文件
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	//创建一个带缓冲的*Reader(输入流)，是带缓冲的，默认缓冲大小为4096，4k，适合读取数据量大的文件，减少io操作提高性能
	r := bufio.NewReader(file)
	for {
		str, err := r.ReadString('\n') //一直读到截止符号
		if err == io.EOF {             //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}

	// 第二次就读取不到了，流水，第一次流完了就没了
	// for {
	// 	str, err := r.ReadString('\n') //一直读到截止符号
	// 	if err == io.EOF {             //io.EOF表示文件的末尾
	// 		fmt.Println("..", str)
	// 		break
	// 	}

	// }
}

//常用的文件操作函数和方法

// 1)打开一个文件进行读操作
// os.Open(name string)(*File,error)
// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；
// 对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
// 2)关闭一个文件
// File.Close()
// 关闭文件，使文件不能用于读写。它返回可能出现的错误。

// 3)其他函数和方法案例讲解(读文件)
// (1)读取文件内容并显示在终端(带缓冲区的方式),适合大文件
//  方法: os.Open,file.close,bufio.NewReader(),reader.ReadString方法和函数
// (2)读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)，这种方式适用于文件不大的情况。
//  方法: ioutil.ReadFile

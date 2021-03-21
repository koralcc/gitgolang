package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1)打开一个已经存在的文件将原来的内容并追加abc english
	//1.打开文件 D:/abc.txt
	filePath := "D:/abc.txt"
	//2.使用写和追加的模式打开
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) //0666只在linux/unix系统中有效
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	//准备写入的5句话
	str := "abc english\r\n" //不同编辑器对换行符的处理不同，统一使用\r\n
	//3.写入时，使用带有缓存的*writer(输出流)
	wr := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		wr.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriteString方法时，其实是先写入
	//到缓存的，所以需要调用Flush方法，将缓存的数据正真写入到文件中，否则文件中会没有数据
	//4.
	wr.Flush()
}

//3)打开一个存在的文件，在原来的内容追加abc english

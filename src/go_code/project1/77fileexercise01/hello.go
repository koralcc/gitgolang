package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	srcPath := "D:/timg.jpg"
	dstPath := "D:/timg2.jpg"
	CopyFile(dstPath, srcPath)

}

func CopyFile(dst string, src string) (written int64, err error) {
	//获取输入流
	fileSrc, err := os.Open(src)
	if err != nil {
		fmt.Println("打开源文件错误", err)
		return
	}
	defer fileSrc.Close()
	readSrc := bufio.NewReader(fileSrc)
	//获取输出流
	writeFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开目标文件错误", err)
		return
	}
	defer writeFile.Close()
	writeDst := bufio.NewWriter(writeFile)

	return io.Copy(writeDst, readSrc)

}

//拷贝文件
//将D:文件拷贝到E:盘下；可用io包下的copy函数实现 //func Copy(dst Writer, src Reader) (written int64, err error)

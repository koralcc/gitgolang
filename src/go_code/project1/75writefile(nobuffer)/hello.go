package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将D:/abc.txt 文件内容导入到 D:/kkk.txt
	//1. 首先将 D:/abc.txt 内容读取到内存
	//2. 将读取到的内容写入 D:/kkk.txt
	fileSourcePath := "D:/abc.txt"
	fileDstPath := "D:/kkk.txt"
	content, err := ioutil.ReadFile(fileSourcePath)
	if err != nil {
		//说明读取文件有误
		fmt.Println("read file err", err)
		return
	}

	err = ioutil.WriteFile(fileDstPath, content, 0666)
	if err != nil {
		//说明写入文件有误
		fmt.Println("write file error", err)
		return
	}

}

//使用ioutil，将一个文件的内容写入到另外一个文件

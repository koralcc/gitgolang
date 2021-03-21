package main

import (
	"fmt"
	"os"
)

func main() {
	filePath1 := "D:/test1.txt"
	// filePath2 := "abc.txt"
	// filePath3 := "abc.txt"
	b, _ := PathExists(filePath1)
	fmt.Println(b)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil //文件或者目录不存在
	}
	return false, err //其他错
}

//判断文件是否存在
// golang判断文件或文件夹是否存在的方法为os.Stat()函数返回的错误值进行判断
// 1) 如果返回的错误为nil，说明文件或文件夹存在
// 2) 如果返回的错误类型使用os.IsNotExist()判断为true，说明文件或文件夹不存在
// 3) 如果返回的错误为其他类型，则不确定是否存在

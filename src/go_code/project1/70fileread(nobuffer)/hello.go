package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file)
	// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。
	// 因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
	if err == nil {
		fmt.Println(string(content))
	}
}

// (2)读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)，这种方式适用于文件不大的情况。
//  方法: ioutil.ReadFile

//ioutil.ReadFile,因为没有显式的open，所以也不需要显式的Close()

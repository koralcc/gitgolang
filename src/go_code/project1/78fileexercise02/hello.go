package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var leterCount, numbCount, spaceCount, otherCount int

	srcPath := "D:/test.txt"
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	read := bufio.NewReader(srcFile)

	for {
		str, err := read.ReadString('\n') //以'\n'为读取的结束符，包含两种结果，遇到'\n'和io.EOF文件末尾

		fmt.Println(str)
		fmt.Println("err", err)
		//为了兼顾中文，可以转化为[]rune
		strs := []rune(str)
		for _, v := range strs {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				leterCount++
			case v >= '0' && v <= '9':
				numbCount++
			case v == '\t' || v == ' ':
				spaceCount++
			default:
				otherCount++

			}
		}
		//必须放在最后，不然最后一个读取不到
		if err == io.EOF {
			break
		}
	}

	// //获取一个字节数组
	// p := make([]byte, 4096)
	// _, err = read.Read(p)
	// if err != nil {
	// 	fmt.Println("文件读取错误:", err)
	// }

	// for _, v := range p {
	// 	switch {
	// 	case v >= 'a' && v <= 'z':
	// 		fallthrough
	// 	case v >= 'A' && v <= 'Z':
	// 		leterCount++
	// 	case v >= '0' && v <= '9':
	// 		numbCount++
	// 	case v == '\t' || v == ' ':
	// 		spaceCount++
	// 	default:
	// 		if v != 0 {
	// 			otherCount++
	// 		}

	// 	}
	fmt.Printf("英文数目:%d\n数字数目:%d\n空格数目:%d\n其他数目:%d\n", leterCount, numbCount, spaceCount, otherCount)
}

//统计一个文件中英文，数字，空格，其他字符数量

package main

import "fmt"

func main() {
	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "abcd"
	slice := str[1:3]
	fmt.Println("str的切片操作", slice)

	//[]byte 转string
	//string 转[]byte
	bytestr := []byte(str)
	fmt.Println("[]byte转string", bytestr)
	byteslice := string(bytestr)
	fmt.Println("string转[]byte", byteslice)
}

//string和slice
//1)string底层是一个byte数组，因此string也可以进行切片处理
//2)string与[]byte的相互转化
// 	string -> []byte   byte[]("abc")
// 	[]byte -> string   string(byteSlice)

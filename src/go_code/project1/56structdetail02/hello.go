package main

import "fmt"

func main() {

	var a A
	var b B
	a = A(b) //可以转化的条件相同的字段(名字、个数和类型必须完全一样)
	fmt.Println(a, b)
}

type A struct {
	Num int
}
type B struct {
	Num int
}

//结构体的注意事项和使用细节
// 2)结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)

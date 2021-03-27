package main

import (
	"fmt"
	"unsafe"
)

func main() {
	caller()
}

func caller() {
	var a int64 = 1
	var b int64 = 2
	c := callee(a, b)
	fmt.Println(c)

	var struct1 = struct{}{}
	var int1 = 12
	var chstruct = make(chan struct{}, 10)
	fmt.Println(unsafe.Sizeof(struct1))
	fmt.Printf("%p\n", &struct1)
	fmt.Printf("%p\n", &int1)
	fmt.Println("chstruct size:", cap(chstruct))
	fmt.Printf("%p\n", &chstruct)
}

func callee(a, b int64) int64 {
	c := a + b
	return c
}

//go tool objdump -s  "main\.main" channeldeadlock.exe | findstr CALL   //windows 反编译某个可执行文件的某个方法，并过滤出CALL语句，linux用grp CALL
//可以用来查看golang源代码的调用始末
//go tool compile -N -l -S stackcall.go

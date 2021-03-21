package main

import "fmt"

func main() {
	num1 := 10
	fmt.Printf("num1的类型是:%T,num1的值是:%v,num1的地址是:%p\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型是:%T,num2的值是:%v,num2的地址是:%p\n", num2, num2, &num2)

}

//Golang 内置函数

// Golang设计者为了编程方便,提供了一些函数,这些函数可以直接使用,我们称为Go的内置函数(在builtin包中)
// 1)len:用来求长度,比如:string,array,slice,map,channel
// 2)new:用来分配内存,主要用来分配值类型,比如:int,float32,struct返回的是指针
// 3)make:用来分配内存,主要用来分配引用类型,比如:channel,map,slice.这个后面会有

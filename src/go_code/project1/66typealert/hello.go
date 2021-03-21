package main

import "fmt"

func main() {

	var a interface{}
	var point Point = Point{1, 2}
	a = point //
	//如何将a赋给一个Point变量
	var b Point
	//b = a //可以吗 cannot use a (type interface {}) as type Point in assignment: need type assertion
	b = a.(Point) //a.(Point)就是类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错
	fmt.Println(b)

	//类型断言的其他案例
	var x interface{}
	var c float32 = 1.1
	x = c
	//y := x //不定义y的类型，使用类型推导也可以直接赋值
	var y float32
	//y := x.(float64) //这里float32必须与x的类型一致,panic: interface conversion: interface {} is float32, not float64
	y = x.(float32)
	fmt.Printf("y类型%T\n", y)

	//类型断言(待检测的)
	var d interface{}
	var e float32 = 1.2
	d = e
	if f, ok := d.(float64); ok {
		fmt.Println("f 转换成功")
		fmt.Printf("转换后的f类型为%T\n", f)
	} else {
		fmt.Println("f 转换失败")
	}

}

type Point struct {
	x int
	y int
}

//类型断言
//需求：如何将一个接口变量，赋给自定义类型的变量 =>引出类型断言

//基本介绍
//类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言

//代码说明：
//在进行类型断言时，如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型

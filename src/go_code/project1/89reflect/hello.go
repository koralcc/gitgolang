package main

import (
	"fmt"
	"reflect"
)

func main() {
	//基本数据类型的反射三个类型的转化
	var num int = 100
	reflectTest01(num)
	//自定义类型反射的三个类型的转化
	stu := Stu{
		Name: "zhangsan",
		Age:  12,
	}
	reflectTest02(stu)

}

func reflectTest01(b interface{}) {
	//通过反射获取传入变量的type，kind，值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal", rVal)
	fmt.Printf("rVal=%v rval Type=%T\n", rVal, rVal)
	//针对不同的类型，使用不同的方法获取值
	n2 := rVal.Int() + 2
	fmt.Println("n2=", n2)
	//3.将rVal转成interface{}
	iV := rVal.Interface()
	//4.将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2:", num2)
}

func reflectTest02(b interface{}) {
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	stu1, ok := iV.(Stu)
	if ok {
		fmt.Println("stu1.Name=", stu1.Name)
	}
}

type Stu struct {
	Name string
	Age  int
}

//反射
// 两个场景
// 1.结构体的tag标签
// 2.桥连接(适配器函数)
// 基本介绍
// 1)反射可以在运行时动态获取变量的各种信息，比如变量的类型(type),类别(kind)
// 2)如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
// 3)通过反射，可以修改变量的值，可以调用关联的方法
// 4)使用反射，需要import("reflect")

//反射重要的函数和概念
// 1)reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型，reflect.Type类型是一个接口类型
// 2)reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型，reflect.Value是一个结构体类型，可以获取到关于该变量的很多信息
// 3)变量，interface{}和reflect.Value是可以相互转化的，这点在实际开发中，会经常使用到

//反射的使用注意细节
// 1)reflect.Value.Kind()，获取变量的类别，返回是一个常量
// 2)Type是类型，Kind是类别，可能是相同的，也可能是不同的
// 比如：var num int = 10  num的Type是int，Kind也是int
// 比如：var stu Student stu的Type是pkg.Student,Kind是struct
// 3)通过反射可以让变量在interface{}和reflect.Value之间相互转换，这点在案例中已经体现
// 4)使用反射的方式来获取变量的值(并返回对应的类型)，要求数据类型匹配，比如x是int，那么就应该使用reflect.Value(x).Int(),而不能使用其他的，否则报panic

//常量补充讲解
// 1)常量使用const修饰
// 2)常量在定义的时候，必须初始化
// 3)常量不能修改
// 4)常量只能修饰bool、数值类型(int、float系列)、string类型
// 5)语法 const identify [type] = Value
// const b = 9/3 //ok
// const c = num/3 //false，常量（也称编译时常量）需要在编译阶段就知道值的类型，用变量计算则成了运行时
// const d = getValue //false
// 6)多个常量的写法
// (1)
// 	const{
// 		a = 1
// 		b = 2
// 	}
// (2)专业写法，一行递增一次
// const{
// 	a = iota  //0
// 	b		  //1
// 	c,d,e     //2
// }
// 表示给a赋值为0
// b在a的基础上+1
// c,d,e在b的基础上+1
// 7)Golang中没有常量名必须大写的规范
// 8)仍然可以通过大小写进行常量的访问限制

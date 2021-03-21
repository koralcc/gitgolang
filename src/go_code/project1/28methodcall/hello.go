package main

import "fmt"

//自定义type
type myint int
type myfuncT func(int, int) int

func main() {
	n := 10
	fmt.Println("main n=", n)
	n2 := 5
	test2(&n2)
	fmt.Println(n2) //&取地址

	//函数也是一种数据类型
	var a int = 10
	var b int = 20
	m := getSum
	fmt.Printf("a的类型为%T，getSUm的类型为%T", m, getSum) //func(int, int) int
	res := m(a, b)
	fmt.Println("方法类型调用结果:", res)
	//函数作为参数传入
	res = myfunc(m, 20, 50)
	fmt.Println("函数作为参数传入结果:", res)
	//自定义type
	var num3 myint = 20
	var num4 int
	num4 = int(num3) //Go认为myint与int是两个不同的类型，需要转化
	fmt.Println(num4)

	res = myfunc(m, 20, 50)
	fmt.Println("函数作为参数传入结果:", res)

	//可变参数
	fmt.Println(sum(1, 2, 3, 4))
	//交换a，b两个值
	var a1, b1 int = 6, 7
	swap(&a1, &b1)
	fmt.Println(a1, b1)
}

func test(n int) {
	n = n + 1
	fmt.Println("test n=", n)
}

//修改基础类型值
func test2(n *int) {
	*n += 2 //*取指针类型变量指向的真正值
}

func getSum(a int, b int) int {
	return a + b
}

//函数式编程
func myfunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//方法作为类型
func myfunc2(funvar myfuncT, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//可变参数
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

//交换两个变量的值
func swap(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}

// func test2(n *int, n2 int) {
// 	*n += 2 //*取指针类型变量指向的真正值
// }

// 函数调用
//1)重点通过画图了解方法调用在内存里面的变化
// (1)在调用一个函数时，会给该函数分配一个新的空间，编译器h会通过自身的处理让这个新的空间和其他栈空间区分开来;
// (2)在每个函数对应的栈中，数据空间是独立的，不会混淆;
// (3)当一个函数执行完毕后，程序会销毁函数对应的栈空间。

//函数注意事项和细节
// 1)函数的形参列表可以是多个，返回值列表也可以是多个
// 2)形参列表和返回值列表的数据类型可以是值类型和引用类型
// 3)函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其他包文件引用，类似public，首字母小写，只能被本包文件使用，其他包文件不能使用，类似private
// 4)函数中的变量是局部的，函数外不生效
// 5)基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改不会影响到原值
// 6)如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&,函数内以指针的方式操作变量，从效果上看类似引用
// 7)Go函数不支持重载  :同一个包里面不能有同名的俩个函数(即使参数不同)
// 8)Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
// 9)函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用(函数式编程)
// 10)为了简化数据类型定义，Go支持自定义数据类型
// 基本语法：type 自定义数据类型名 数据类型 //理解：相当于一个别名
// 案例：type myint int // 这时，myint就等价于int来使用了
// 案例：type mySum func(int,int)int//这时，mysum就等价于一个函数类型func(int,int)int
//  11)支持对返回值的命名，这样就不需要在意返回的顺序，直接return即可
// 12)使用_标识符，忽略返回值
// 13)Go支持可变参数
// //支持0到多个参数
// func sum(args...int)sum int{

// }
// //支持1到多个参数
// func sum(n1 int,args... int)sum int{

// }
// (1)args是slice切片，通过args[index]访问
// (2)如果一个函数的形参列表有可变参数，则可变参数需放在形参列表的最后

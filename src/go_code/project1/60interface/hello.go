package main

import "fmt"

func main() {
	com := Computer{}
	pho := Phone{}
	cam := Camera{}

	com.Working(pho)
	com.Working(cam)
	//接口是一个指针类型
	stu := Stu{}
	stu.Name = "jimmy"

	fmt.Println(stu.Name)

	//接口的继承
	var e BInterface = E{}
	e.Test01()
}

//Usb:引入一个usb插camera和手机的接口
//Usb: to insert or convert power
type Usb interface {
	Start() //接口里面不能有变量，只能有没实现的方法
	Stop()
}

type Computer struct {
}

type Camera struct {
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作......")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作......")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作......")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作......")
}

//编写一个working方法，接收一个Usb接口变量，只要是实现了Usb接口(所谓实现Usb接口，就是指实现了Usb接口的所有方法)
func (c Computer) Working(usb Usb) {
	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu Say()")
}

func (stu Stu) setStuName() {
	stu.Name = "Hello"
}

type BInterface interface {
	Test01()
	Test02()
}

type CInterface interface {
	Test01()
	Test03()
}

type DInterface interface {
	BInterface
	CInterface
}

type E struct {
}

func (e E) Test01() {
	fmt.Println("e 的test01方法")
}

func (e E) Test02() {
	fmt.Println("e 的test02方法")
}

func (e E) Test03() {
	fmt.Println("e 的test03方法")
}

//接口
//interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。
//到某个自定义类型(比如结构体Phone)要使用的时候，再根据具体情况把这些方法写出来(实现)
// 1)基本语法
// type 接口名 interface{
// 	method1(参数列表)返回值列表
// 	method2(参数列表)返回值列表
// }

// 实现接口所有方法
// func(t 自定义类型)method1(参数列表)返回值列表{
// 	//方法实现
// }
// func(t 自定义类型)method2(参数列表)返回值列表{
// 	//方法实现
// }
// 说明：(1)接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态
// 和高内聚低耦合的思想
// 	     (2)Golang中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量
// 就实现这个接口。因此，Golang中没有implement这样的关键字

// 2)注意事项和细节
// (1)接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
// (2)接口中所有的方法都没有方法体，即都是没有实现的方法
// (3)在Golang中，一个自定义类型需要将某个接口的所有	方法都实现，我们说这个自定义类型实现了该接口。
// (4)一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
// (5)只要自定义数据类型，就可以实现接口，不仅仅是结构体类型。例如：type integer int
// (6)一个自定义类型可以实现多个接口
// (7)Golang接口不能有任何变量
// (8)一个接口(比如A接口)可以继承多个别的接口(比如B，C接口)，这时如果要实现A接口，也必须将B，C接口的方法也全部实现
// (9)interface 类型默认是一个指针(引用类型)，如果没有对interface初始化就使用，那么会输出nil
// (10)空接口interface{}没有任何方法，所以所有类型都实现了空接口，即我们可以把任何一个变量赋给空接口

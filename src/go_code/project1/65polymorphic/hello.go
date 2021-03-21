package main

import "fmt"

func main() {
	var usbs [2]Usb
	usbs[0] = Phone{}
	usbs[1] = Computer{}
}

type Usb interface {
	charg()
}

type Phone struct {
}

type Computer struct {
}

func (p Phone) charg() {
	fmt.Println("手机充电中")
}

func (c Computer) charg() {
	fmt.Println("手机充电中")
}

//多态，Golang的多态通过接口体现
// >> 多态参数
// 在前面的Usb接口案例中，Usb usb ，即可以接收手机变量，又可以接收相机变量，就体现了Usb接口的多态
// >> 多态数组
// 相当于实现了将不同的数据类型放到了同一个数组中

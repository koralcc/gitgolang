package main

import "fmt"

func main() {
	//定义一个多态数组，将phone和camera分别放进去
	us := [3]Usb{
		Phone{"小米"},
		Phone{"苹果"},
		Camera{"尼康"},
	}
	for _, v := range us {
		v.start()
		v.charge()
		if p, ok := v.(Phone); ok {
			p.call()
		}
		v.end()
	}

	//利用v.(type)获取类型
	var a int8 = 1
	var b int = 23
	var c bool = true
	var d float64 = 64
	var e []string
	var s1 Student = Student{}
	var s2 *Student = &Student{}
	TypeJudge(a, b, c, d, e, s1, s2)
	//t := e.(type) //use of .(type) outside type switch,.(type)只能在switch中使用
}

type Computer struct{}
type Phone struct {
	Name string
}
type Camera struct {
	Name string
}

type Usb interface {
	start()
	charge()
	end()
}

type Student struct{}

func (c Computer) working(u Usb) {
	u.start()
	u.charge()
	u.end()
}

func (p Phone) start() {
	fmt.Printf("%s 开始充电\n", p.Name)
}

func (p Phone) charge() {
	fmt.Println("Phone充电中.......")
}

func (p Phone) end() {
	fmt.Printf("Phne%s 结束充电\n", p.Name)
}

func (p Phone) call() {
	fmt.Printf("%s 打电话\n", p.Name)
}

func (p Camera) start() {
	fmt.Printf("%s 开始充电\n", p.Name)
}

func (c Camera) charge() {
	fmt.Println("Camera充电中......")
}

func (p Camera) end() {
	fmt.Printf("%s 结束充电\n", p.Name)
}

//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(source ...interface{}) {
	//source是一个切片
	for i, v := range source {
		switch v.(type) { //注意这里type是一个关键字，固定写法
		case bool:
			fmt.Printf("第%d个参数是bool类型，值是%v\n", i, v)

		case float32, float64:
			fmt.Printf("第%d个参数是float类型，值是%v\n", i, v)

		case int, int16, int32, int64, int8:
			fmt.Printf("第%d个参数是int类型，值是%v\n", i, v)

		case string:
			fmt.Printf("第%d个参数是string类型，值是%v\n", i, v)

		case nil:
			fmt.Printf("第%d个参数是nil类型，值是%v\n", i, v)
		case Student:
			fmt.Printf("第%d个参数是Student类型，值是%v\n", i, v)
		case *Student:
			fmt.Printf("第%d个参数是*Student类型，值是%v\n", i, v)
		default:
			fmt.Printf("第%d个参数不知道是什么类型", i)

		}
	}
}

//类型断言案例
//电脑，手机，摄像机，结合多态数组操作

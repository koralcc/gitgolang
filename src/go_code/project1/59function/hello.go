package main

import "fmt"

func main() {
	var b integer = 2
	//(&b).addI() 与b.addI()相同
	b.addI() //会自动转换，是否转化成&b与方法addI的方法类型有关
	fmt.Println(b)

	var stu1 = Student{"zhangsan", 18}
	fmt.Println("stu1", &stu1)
	fmt.Printf("stu1本身值为%v,stu1地址为%p\n", &stu1, &stu1)

	var int1 int = 6

	fmt.Println("int1", fmt.Sprintf("%d", int1))
}

type integer int

//cannot define new methods on non-local type int
//go语言不允许为简单的内置类型添加方法。
// func (a *int) add() {
// 	*a++
// }

func (b *integer) addI() {
	*b++
}

type Student struct {
	Name string
	age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v],age=[%v]", stu.Name, stu.age)
	return str
}

//方法
// 1)在某些情况下，我们需要声明(定义)方法，比如Person结构体；除了有一些字段外，person结构体还有一些行为，比如，学习，唱歌等
// 2)Golang中的方法是作用在指定的数据类型上的(即和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅仅是struct，比如：type integer int
// 3)方法的调用和传参机制
//方法的掉用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当作实参传递给方法
//    4)方法的声明(定义)
//    func (receive type) methodName(参数列表) (返回值列表){
// 	   方法体
// 	   return 返回值
//    }
//    (1)参数列表，标识方法输入
//    (2)receive type: 表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
//    (3)receive type: 可以是结构体，也可以是其他自定义类型
//    (4)receive:就是type类型的一个变量(实例)，比如：Person结构体的一个变量(实例)
//    (5)参数列表：表示方法输入
//    (6)返回值列表：表示返回值，可以多个
//    (7)方法主体：表示为了实现某一功能代码块
//    (8)return 语句不是必须的

//方法注意事项和细节讨论
// 1)结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
// 2)如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
// 3)Golang中的方法作用在指定的数据类型上的，因此自定义类型都可以有方法，而不仅仅是struct，比如int,float等都可以有方法
// 4)方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其他包访问
// 5)如果一个变量实现了String()这个方法，那么fmt.Println()默认会调用这个变量的String()进行输出

//方法和函数的区别
// 1)调用方式不一样
// 函数的调用方式：函数名(实参列表)
// 方法的调用方法：变量.方法名(实参列表)
// 2）对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，接收者为指针类型时，不能将值类型的数据直接传递
// 3)对于方法(比如struct的方法),接收者为值类型时，可以直接用指针类型的变量调用方法，反过来也可以，就是无论接收类型为什么
// 都可以用指针/值类型变量调用，起到决定作用的是接收类型

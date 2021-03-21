package main

import (
	"fmt"
	"go_code/project1/63extends/model"
)

func main() {
	pupil := model.Pupil{}
	fmt.Println(pupil.Student.Name)
	//fmt.Println(pupil.Name)
	pupil.Student.Name = "学生"
	pupil.Person.Name = "人类"
	//fmt.Println("pupil Name", pupil.Name)
	//fmt.Println(pupil.Student.age)，不同包不能使用继承的小写变量和方法

	//多重继承具有相同属性时
	c := C{}
	c.A.Name = "tom"
	c.Age = 18
	c.A.Score = 100
	c.A.Weight = 90
	//fmt.Println(c.Name) //两个匿名结构体都有Name，但是本身结构体没有时报错：ambiguous(模棱两可的) selector c.Name
	fmt.Println("c.Age:", c.Age)      //c自己的age
	fmt.Println("c.A.Age:", c.A.Age)  //c.A自己的age
	fmt.Println("c.B.ge:", c.B.Age)   //c.B自己的age，本身结构体和继承的存在相同的属性/方法时，按指定的取
	fmt.Println("c.Score", c.Score)   //c自身的score
	fmt.Println("c.Weight", c.Weight) //先找c没有，找A找到了
	c.SayHello()                      //**注意这里虽然是c调用的sayhello方法，但是实际这个方法是属于A的，所以会获取到A的Name
	//匿名结构体赋值,也可以在初始化的时候进行赋值

	c2 := C{
		A: A{
			Name:   "A",
			Age:    20,
			Score:  20,
			Weight: 29,
		},
		B: B{
			Name: "B",
			Age:  21,
		},
		Age:   23,
		Score: 23,
	}
	//或者不指定属性按顺序初始化，注意此时必须保持一致，一不写属性全部不写(包括里面的和外面的)
	//c2 := C{A{"A", 20, 20, 29}, B{"B", 21}, 23, 23}
	fmt.Println(c2)
}

type A struct {
	Name   string
	Age    float64
	Score  float64
	Weight float64
}

type B struct {
	Name string
	Age  float64
}

type C struct {
	A
	B
	Age   float64
	Score float64
}

func (a A) SayOk() {
	fmt.Println("a say ok", a.Name)
}

func (b B) SayOk() {
	fmt.Println("b say ok", b.Name)
}

func (a A) SayHello() {
	fmt.Println("a say hello", a.Name)
}

// //继承可以解决代码复用的问题，让我们的编程更加接近人类思维
// 当多个结构体存在相同的属性(字段)和方法时，可以从这些结构体中抽象出结构体(比如student)，在该结构体中定义这些相同的属性和方法
// 其他的结构体不需要重新定义这些属性(字段)和方法，只需嵌套一个Student匿名结构体即可
// 也就是说，在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性

// //继承的细节
// 1)结构体可以使用嵌套匿名结构体的所有字段和方法(同一个包中)，即：首字母大写或者小写的字段、方法，都可以使用
// 2)匿名结构体字段访问可以简化，即，直接访问(只有一个继承)
// 当我们直接通过pup访问字段或者方法时，其执行流程如下，如pup.Name
// 编译器会先看pup对应的类型有没有Name，如果有，则直接调用pup类型的Name字段
// 如果没有就去看pup中嵌入的匿名结构体Student中有没有Name字段，如果有就调用,如果都找不到就报错
// 如果有两个或者以上继承时，两个父结构体都有相同的字段或者方法但是自shiyong身结构体没有相同的字段或方法时，会报错 ambiguous selector pupil.Name（第四点）
// 3)当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
// 4)结构体嵌入两个(或多个)匿名结构体，如果两个匿名结构体有相同的字段或方法(同时结构体本身没有相同的字段和方法)，在直接访问(pupil.Name)时，就必须明确指定匿名结构体名字，
// 否则编译报错ambiguous selector pupil.Name
// 5)如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体字段或方法时，必须带上结构体的名字
// 6)嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值

// 自己总结(精髓)：## 结构体与匿名结构体有相同的属性/方法时，直接取取的是结构体的属性/方法
// 				  ## 结构体与匿名结构体没有相同的属性/方法时，直接取，如果只有一个匿名结构体存在，可以取到，如果有多个匿名结构体存在，则报错
//				  ###############################################################################

//多重继承说明
// 如果一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而试先了多重继承
// 1)如果嵌入的匿名结构体有相同的字段名或方法，则在访问时，需要通过匿名结构体类型来区分
// 2)为了保证代码的简洁性，尽量不使用多重继承

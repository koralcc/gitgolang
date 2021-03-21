package main

import "fmt"

func main() {

	//声明结构体变量
	//方式二
	p2 := Person{"zhangsan", 18, [5]float64{11, 12, 13, 12, 11}, new(int), make([]int, 3), make(map[string]string)}
	fmt.Println(p2)

	//方式三
	var p3 *Person = new(Person)
	//因为p3是一个指针，因此标准的给字段的赋值方式为
	(*p3).Name = "tom" //一定要括起来的原因是因为 .的优先级>*
	//等价于如下的写法，go的设计者为了程序员使用方便，底层对p3.Name进行了处理，加上了取值运算*p
	p3.Name = "tom"

	//方式四
	var p4 *Person = &Person{}
	(*p4).Name = "jimmy"

	//方式一
	var cat1 Cat //值类型，声明就分配内存，且cat1的内存就为第一个属性Name的内存
	fmt.Printf("definecat1的地址为%p,cat1.Name的地址为%p,cat1.Age的地址为%p,值为%v\n", &cat1, &cat1.Name, &cat1.Age, cat1)
	//cat1的地址为0xc00001c080,cat1.Name的地址为0xc00001c080,cat1.Age的地址为0xc00001c090,值为{ 0  }
	cat1.Name = "audi"
	cat1.Age = 0
	cat1.Color = "white"
	cat1.Hobby = "eating"
	fmt.Printf("aftercat1的地址为%p,cat1.Name的地址为%p,cat1.Age的地址为%p,值为%v\n", &cat1, &cat1.Name, &cat1.Age, cat1)

	cat2 := cat1 //这个是底层复制了一份出来cat2再重新指向
	fmt.Printf("cat2的地址为%p,cat2.Name的地址为%p,cat2.Age的地址为%p,值为%v\n", &cat2, &cat2.Name, &cat2.Age, cat2)

	var te1 test
	fmt.Printf("%p,%p,%p\n", &te1, &te1.name, &te1.age)

	//结构体的初始化
	var p1 Person
	p1.Slice = []int{1, 2, 3}
	fmt.Printf("name 地址%p,age 地址 %p,score地址%p,ptr 地址%p,Slice 地址%p,Map1 地址%p\n", &p1.Name, &p1.Age, &p1.Score, &p1.Ptr, &p1.Slice, &p1.Map1)
	//name 地址0xc00010a000,age 地址 0xc00010a010,score地址0xc00010a018,ptr 地址0xc00010a040,Slice 地址0xc00010a048,Map1 地址0xc00010a060
	//这里的地址是ptr本身的地址，它存的值是nil,即指向的值为nil，没分配内存空间初始化的时候，指针，切片，映射等初始值都是nil，即还没有分配空间，还不能使用
	fmt.Printf("name 值%v,ptr 值%p,Slice 值%p,Slice[0]的地址为%p\n", p1.Name, p1.Ptr, p1.Slice, &p1.Slice[0])
	//用%v取引用类型，直接取出的是指向的真正值，p1.Slice的值是指针类型，所以需要%p进行格式化，而不是%v
	//name 值,ptr 值0x0,Slice 值0xc000012400,Slice[0]的地址为0xc000012400

}

//Cat struct
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

type test struct {
	name int
	age  int
}

//Person str
type Person struct {
	Name  string
	Age   int
	Score [5]float64
	Ptr   *int              //指针
	Slice []int             //切片
	Map1  map[string]string //映射
}

//struct 结构体
// 结构体和结构体变量(实例)的区别和联系
// 1)结构体是自定义的数据类型，代表一类事物
// 2)结构体变量(实例)是具体的，实际的，代表一个具体变量
// 3）结构体变量在内存中的布局
// 4)结构体的声明:
// 		type 标识符 struct {
// 			filed1 type
// 			filed2 type
// 		}
//   首字母大写，结构体类型可以被其他包引用，字段大写可以被其他包引用，小写只能本包使用
//   思考一个问题：本包引用其他包方法，在本包调用私有的结构体，能成功？答案：不能
// 5)字段/属性
// (1)从概念或叫法上看：结构体字段 = 属性 = filed
// (2)字段是结构体的一个组成部分，一般是基本数据类型、数组，也可以是引用类型，比如我们前面定义Cat结构体的name就是属性

// 创建结构体变量和访问结构体字段
// 1)方式一-直接声明
// 案例演示：var person Person
// 2)方式二：
// var person Person = Person{}
// 3)方式三:
// var person *Person = new(Person)
// 4)方式四：
// var person *Person = &Person{}
// (1)第三种和第四种方式返回的是结构体指针
// (2)结构体指针访问字段的标准方式应该是:(*结构体指针).字段名，比如:(*person).Name = "tom"
// (3)但go做了一个简化，也支持结构体指针.字段名，比如，person.Name = "tome"。更加符合程序员
// 使用习惯，go编译器底层对person.Name做了转化(*person).Name

// 字段/属性注意事项和细节
// 1)字段声明语法同变量，示例：字段名 字段类型
// 2)字段的类型可以为：基本类型、数组或引用类型
// 3)在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前面的一样：
// 布尔类型：false
// 数值：0
// 字符串：""
// 数组类型和它的元素类型有关，比如socre [3]int 则为[0,0,0]
// 指针、slice和map的零值都是nil，即还没有分配空间
// 4)不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个

//结构体的注意事项和使用细节(详情见下面几节)
// 1)结构体的所有字段在内存中都是连续的,指针类型也是
// 2)结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)
// 3)结构体进行type重新定义(相当于取别名)，Golang认为是新的数据类型，但是相互间可以强转
// 4)struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化

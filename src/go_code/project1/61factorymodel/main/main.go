package main

import (
	"fmt"
	"go_code/project1/61factorymodel/model"
)

func main() {
	//大写的直接引用即可
	stu := model.Student{
		Name: "tom",
		Age:  12,
	}
	fmt.Println("stu:", stu)

	//小写的直接引用报错 cannot refer to unexported name model.student
	// stu1 := model.student{
	// 	Name: "tom",
	// 	Age:  12,
	// }
	//var stu1 model.student = model.Newstudent("jimmy", 18) cannot refer to unexported name model.student
	stu1 := model.Newstudent("jimmy", 18) //必须这样去实例化 不能用var类型的形式去做
	fmt.Println("stu1", stu1)
	fmt.Println("stu1.age", stu1.GetAge())
}

//工厂模式
// 对于首字母小写的结构体，使用工厂模式进行实例化；对于小写的属性，亦可使用
// 工厂模式，提供一个大写的方法，返回一个实例/变量

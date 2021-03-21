package main

import "fmt"

func main() {
	var stu1 Student
	var stu2 Stu
	//stu2 = stu1 //cannot use stu1 (type Student) as type Stu in assignment
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)

	var int1 int
	var int2 integer
	int2 = int(int1)
	fmt.Println(int1, int2)
}

type Student struct {
	Name string
	age  int
}

type Stu Student
type integer int

//结构体的注意事项和使用细节(详情见下面几节)
// 3)结构体进行type重新定义(相当于取别名)，Golang认为是新的数据类型，但是相互间可以强转

package main

import "fmt"

func main() {
	var name string
	var age byte
	var salary float32
	var ispass bool
	//方式一
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&salary)
	// fmt.Println("是否通过考试：")
	// fmt.Scanln(&ispass)
	// fmt.Printf("%v的年龄是%v\n，薪水是%f\n，考试通过情况是%v", name, age, salary, ispass)

	//方式二:按对应的格式输入
	fmt.Println("请输入你的姓名，年龄，薪水，考试通过情况，并用空格隔开")
	fmt.Scanf("%s%d%f%t", &name, &age, &salary, &ispass)
	fmt.Scanf("%s%d%f%t", &name, &age, &salary, &ispass)
	fmt.Printf("%v的年龄是%v\n，薪水是%f\n，考试通过情况是%v", name, age, salary, ispass)
}

func add(a int) {
	a++
}

//键盘输入语句
// 在编程中，需要接收用户输入的数据，就可以使用键盘输入语句来获取
// 1)fmt.Scanln
// 2)fmt.Scanf

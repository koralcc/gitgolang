package model

import "fmt"

type Person struct {
	Name   string
	Height float64
}

type Student struct {
	Name string
	age  float64
}

type Pupil struct {
	stu Student //组合，有名结构体
	Student
	//Name string
	Person
}

func (stu *Student) SayOk() {
	fmt.Println("Student say ok", stu.Name)
}

func (stu *Student) sayHello() {
	fmt.Println("Student say hello", stu.age)
}

//
// func (pup *Pupil) SayOk() {
// 	fmt.Println("Student say ok", pup.Name)
// }

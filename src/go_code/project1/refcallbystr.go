package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stu Student
	ref := reflect.ValueOf(stu)

	f := ref.MethodByName("SayHello")
	f.Call([]reflect.Value{})
}

type Student struct {
}

func (stu Student) SayHello() {
	fmt.Println("hello workd")
}

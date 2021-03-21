package util

import "fmt"

type person struct {
	name string
	age  float64
}

type Student struct {
	*person
}

func (p person) sayOk() {
	fmt.Println(p.name)
}

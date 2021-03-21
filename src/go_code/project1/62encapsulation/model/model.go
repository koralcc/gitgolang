package model

import "fmt"

type person struct {
	Name   string
	age    float64
	salary float64
}

func NewPerson(n string) *person {
	return &person{
		Name: n,
	}
}

func (per *person) SetAge(a float64) {
	if a < 0 || a > 200 {
		fmt.Println("年龄范围不正确......")
		return
	}
	per.age = a
}

func (per *person) GetAge() float64 {
	return per.age
}

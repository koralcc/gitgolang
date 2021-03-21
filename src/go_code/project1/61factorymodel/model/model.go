package model

type Student struct {
	Name string
	Age  float64
}

type student struct {
	Name string
	age  float64 //
}

//工厂模式实例化
func Newstudent(n string, a float64) student {
	return student{
		Name: n,
		age:  a,
	}
}

//提供对age 私有变量的访问方法
func (stu student) GetAge() float64 {
	return stu.age
}

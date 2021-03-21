package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 string = "12"
	var f1 string = "12.2387878873"
	var b1 string = "true"
	var zero string = "hello"
	//var c1 string = "A"

	ii1, _ := strconv.ParseInt(i1, 10, 64)
	//i1：字符串，10：期望转化的进制，0：期望转化的类型
	fmt.Printf("ii1type=%T,ii=%v\n", ii1, ii1)

	ff1, _ := strconv.ParseFloat(f1, 32)
	fmt.Printf("ff1type=%T,ff1=%v\n", ff1, ff1)
	//parseFloat:返回的是float64，此时的32是指期望的类型，如果此数超过32位的精度，可能造成精度缺失，可以修改成64确保精度准确

	bb1, _ := strconv.ParseBool(b1)
	fmt.Printf("bb1type=%T,bb1=%v\n", bb1, bb1)

	zeroo, _ := strconv.ParseInt(zero, 10, 64)
	fmt.Println("zeroo=", zeroo)

	bb2, _ := strconv.ParseBool(zero)
	fmt.Println("zeroo=", bb2)

}

//string转基本类型
// strconv.ParseInt
// strconv.ParseFloat
// strconv.ParseBool

// 注意事项：string转基本类型，要确保String类型能转成有效的数据，比如我们可以把'123'转化成一个整数，但是不能把'hello'转化成一个整数，如果这样做
// Golang会直接将其转成0
//*在字符串转基本类型失败时(如上注意事项)，不会报错，会得到对应类型的0值*

package utils

import "fmt"

//根据操作符进行计算
func Calc(n1 float32, n2 float32, op string) float32 {

	fmt.Println(I) //相同包变量和方法直接引用
	switch op {
	case "+":
		return n1 + n2
	case "2":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		return 0
	}
}

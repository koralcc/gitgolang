package main

import "fmt"

//编写一个函数,输出100以内的所有素数(只能被1和本身整除,且大于1),每行显示5个;并求和
func main() {
	var count int //打印次数
	for i := 2; i < 101; i++ {

		if isSs(int64(i)) {
			continue
		}
		count++
		fmt.Printf("%v ", i)

		if count >= 5 {
			count = 0
			fmt.Println()
		}
	}
}

func isSs(a int64) bool {

	for i := 2; int64(i) < a; i++ {
		if a%int64(i) == 0 {
			return true
		}
	}

	return false
}

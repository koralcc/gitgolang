package main

import "fmt"

func main() {
	fmt.Println(clumsy(6))
}

func clumsy(N int) int {
	stack := []int{N}
	var index = 0 //控制*/+-
	N--
	for N > 0 {
		switch index % 4 {
		case 0: //*
			stack[len(stack)-1] *= N
		case 1: ///
			stack[len(stack)-1] /= N
		case 2: //+
			stack = append(stack, N)
		default: //-
			stack = append(stack, -N)

		}
		N--
		index++
	}
	var ret int
	for _, v := range stack {
		ret += v
	}
	return ret
}

//【1006】笨阶乘

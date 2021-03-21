package main

import "fmt"

func main() {
	//1打印空心金字塔，每层 2n-1
	//实心金字塔
	// 	   *           3空格
	//    ***          2
	//   *****         1
	//  *******        0
	var total int = 10
	for i := 0; i < total; i++ {
		//空格
		for k := 0; k < total-i; k++ {
			fmt.Print(" ")
		}
		//星星
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//空心金字塔，在实心金字塔的基础上去掉非最后一行的中间星星
	// 	   *
	//    * *
	//   *   *
	//  *******
	var total2 int = 10
	for i := 0; i < total2; i++ {
		//空格
		for k := 0; k < total2-i; k++ {
			fmt.Print(" ")
		}
		//星星
		for j := 0; j < 2*i-1; j++ {
			//只有最后一行和每行的第一个和最后一个才打印
			//if i == total2-1 || (i != total2-1 && (j == 0 || j == 2*i-2)) {
			if i == total2-1 || j == 0 || j == 2*i-2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	//2打印99乘法表
	// 1 * 1 = 1
	// 1 * 2 = 2	2 * 2 = 4
	// 1 * 3 = 3	2 * 3 = 6	3 * 3 = 9
	// 1 * 4 = 4	2 * 4 = 8	3 * 4 = 12	4 * 4 = 16
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

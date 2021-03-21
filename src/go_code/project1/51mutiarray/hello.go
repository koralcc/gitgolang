package main

import "fmt"

func main() {
	// 000000
	// 000100
	// 001000
	var mutiarry [3][6]int
	mutiarry[1][3] = 1
	mutiarry[2][2] = 1

	fmt.Println("mutiarry", mutiarry)
	//按要求输出
	for i := 0; i < 3; i++ {
		for j := 0; j < len(mutiarry[i]); j++ {
			fmt.Print(mutiarry[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("mutiarry的地址为:%p", &mutiarry)
	fmt.Printf("mutiarry[0]的地址%p,mutiarry[1]的地址%p,mutiarry[2]的地址%p\n", &mutiarry[0], &mutiarry[1], &mutiarry[2])
	//二维数组，第一维的取值分别存储三个指针指向第二维的三个数组首地址
	fmt.Printf("mutiarry[0][0]的地址为%p,mutiarry[1][0]的地址为%p,mutiarry[2][0]的地址为%p\n", &mutiarry[0][0], &mutiarry[1][0], &mutiarry[2][0])
	//且每个地址之间相隔64个字节，因为二维总共有6个元素，一个int64是8个字节 6*8 =48
}

// 二维数组
//一维数组里面的元素是一个数组这样就构成了二维数组
// 000000
// 000100
// 001000
// 1)语法 var 数组名[大小][大小]类型
//    2)比如var mutiarry [3][6]int，再赋值
//    3)二维数组在内存中的布局(非常重要)

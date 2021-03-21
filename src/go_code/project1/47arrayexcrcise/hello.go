package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1创建一个26个元素得数组,分别放置'A'-'Z'
	var wordArr [26]byte

	for i := 0; i < 26; i++ {
		wordArr[i] = 'A' + byte(i)
	}

	for i, v := range wordArr {
		fmt.Printf("wordArr[%d]=%c\n", i, v)
	}
	//2.求出一个数组得最大值,并得到对应得下标
	var wordArr2 [5]int = [...]int{1, 8, 87, 67, 56}
	var max int = wordArr2[0]
	var maxindex int
	for i, v := range wordArr2 {
		if v > max {
			max = v
			maxindex = i
		}
	}
	fmt.Printf("最大值为%v,下标为%v", max, maxindex)

	//3.随机生成5个数,并将其反转打印(有很多实现方法,但是这里我们用交换位置来实现)
	Arr3 := [5]int{1: 1, 3: 3, 2: 2, 4: 4, 0: 5}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		a := rand.Intn(100)
		Arr3[i] = a
	}

	fmt.Println("反转前得arr3:", Arr3)
	// 1 2 3 4 5,五个元素就需要交换5/2 2次   0   4   5-1
	//										1   3   5-2
	var temp int
	for i := 0; i < len(Arr3)/2; i++ {
		temp = Arr3[i]
		Arr3[i] = Arr3[len(Arr3)-1-i]
		Arr3[len(Arr3)-1] = temp
	}

	fmt.Println("反转后得arry为:", Arr3)
	ch := make([]int, 12)
}

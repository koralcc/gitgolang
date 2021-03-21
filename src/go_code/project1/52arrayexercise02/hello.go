package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.随机生成10个整数(1_100的范围)保存到数组，并倒叙打印以及求平均值、最大值，和最大值的下标、并查找里面是否有55
	var array1 [10]int

	rand.Seed(time.Now().Unix()) //正确的用法，srand 一个程序只能调用一次。或者说一个线程只调用一次
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second)
		//rand.Seed(time.Now().Unix()) //rand 是固定的迭代函数，对一个初始值不停嵌套计算，输出每次计算的结果。输出的序列有不错的统计性质，所以叫伪随机数。srand 就是设置这个初始值。不停地设置相同的初始值，rand 就总是初始值经过一次计算的结果，也就不会变。这是错误的用法。正确的用法，srand 一个程序只能调用一次。或者说一个线程只调用一次（对于常见的每个线程一个随机数发生器状态的 C 运行时库）。
		array1[i] = rand.Intn(100) + 1
	}

	fmt.Println("随机生成的数组", array1)

	//获取平均值
	fmt.Println("arry1的平均值为：", getAverageValue(array1))
	//获取最大值以及下标
	maxIndex, maxValue := getMaxvalue(array1)
	fmt.Println("arr1的最大下标和最大值为：", maxIndex, maxValue)
	//获取是否含有55
	isConatains55 := findxValue(array1)
	fmt.Println("#1#arry1是否含有55[-1不含]", isConatains55)

	//2.已知有个排序好(升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
	insertParam := 50
	array2 := [5]int{2, 10, 20, 30, 40}
	array3 := insertParamToArrayAsc(array2, insertParam)
	fmt.Printf("#2#插入%v前的数组：%v,插入后的数组%v", insertParam, array2, array3)

	//3.定义一个三行四列的数组，逐个从键盘输入值，编写程序将四周的数据清0
	// 1111		   0000
	// 1111  ==>>  0110
	// 1111        0000
	// 行数是0或者行数是len(arr(x))的都要清空
	// 列数是0或者列数是len(arr(y))的都要清空
	var arry4 = [3][4]int{{12, 13, 14, 15}, {112, 113, 141, 151}, {122, 133, 124, 125}}
	clear3x4yRound(&arry4)
	fmt.Println("#3#befor clear", arry4)
	fmt.Println("#3#after clear", arry4)

	//4.定义一个4行4列的二维数组，逐个从键盘输入值，然后将第1行和第4行的数据进行交换，
	//将第2行和第3行数据进行交换
	var arry5 = [4][4]int{{2, 2, 3, 4}, {1, 2, 3, 4}, {8, 7, 6, 8}, {2, 2, 2, 2}}
	fmt.Println("#4#替换前", arry5)
	replace44Arr(&arry5)
	fmt.Println("#4#替换后", arry5)
}

//获取平均值
func getAverageValue(arr [10]int) float64 {
	total := 0
	for i := 0; i < 10; i++ {
		total += arr[i]
	}
	return float64(total) / 10
}

//获取最大值，和最大值下标
func getMaxvalue(arr [10]int) (i, v int) {

	var maxValue = arr[0]
	var maxIndex = 0
	//获取最大值
	for i := 1; i < 10; i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
			maxIndex = i
		}
	}
	return maxIndex, maxValue
}

//查找该数组里面是否有55，并返回下标
func findxValue(arr [10]int) int {
	for i, v := range arr {
		if v == 55 {
			return i
		}
	}

	return -1
}

//插入元素到数组里面
func insertParamToArrayAsc(arr [5]int, newInt int) [6]int {
	var arr2 [6]int

	//将原数组元素copy到新加的数组中
	for i := 0; i < len(arr); i++ {
		arr2[i] = arr[i]
	}

	//fmt.Println(" init arr2", arr2)
	//1.找到要插入的位置，然后后续的元素全部往后移动一位
	start := 0
	end := 4
	//直接找中间的元素，如果arr[mid]>newInt 往左找第一个小于newInt的元素,如果没有，则要插入的位置为 [0]
	//如果arr[mid]<newInt 往右找第一个大于newInt的元素,如果没有，则要插入的位置为 [len(arr)]
	mid := (start + end) / 2
	var insertIndex int
	if arr[mid] > newInt {
		//往左边找
		for i := mid - 1; i >= 0; i-- {
			if arr[i] < newInt {
				insertIndex = i + 1
				break //找到第一个小于插入值的下标，直接返回
			}
		}
	} else if arr[mid] < newInt {
		//往右边找
		for i := mid + 1; i < len(arr)-1; i++ {
			if arr[i] > newInt {
				insertIndex = i - 1
				break
			}
		}
		insertIndex = len(arr2) - 1
	} else {
		insertIndex = mid + 1
	}

	if insertIndex != len(arr2)-1 {
		//只要不是在最后面插入元素，都要把该元素后的每一个元素都往后移动一个位置
		for i := len(arr2) - 1; i > insertIndex; i-- {
			arr2[i] = arr2[i-1]
		}
	}
	// fmt.Println("插入点", insertIndex)
	// fmt.Println(" before insert arr2", arr2)

	//最后插入元素
	arr2[insertIndex] = newInt

	return arr2
}

//清空三行四列周边的数据为0
func clear3x4yRound(arr *[3][4]int) {
	// 行数是0或者行数是len(arr(x))的都要清空
	// 列数是0或者列数是len(arr(y))的都要清空
	for i, v := range arr {
		for i2, _ := range v {
			if i == 0 || i == len(arr)-1 || i2 == 0 || i2 == len(v)-1 {
				(*arr)[i][i2] = 0
			}
		}
	}
}

//4*4的二维数组进行替换
func replace44Arr(arr *[4][4]int) {

	var temp int
	for i, v := range arr {
		if i > len(arr)/2-1 {
			break
		}
		for i2, v2 := range v {
			temp = v2
			arr[i][i2] = arr[len(arr)-i-1][i2]
			arr[len(arr)-i-1][i2] = temp
		}
	}
}

package main

import "fmt"

func main() {
	//1.使用冒泡排序进行分析和处理76,58,67,18,0,9,
	var array = [6]int{76, 58, 67, 18, 0, 9}
	bubbleSort(&array)
	fmt.Println("array", array)

	//2.顺序查找
	name := [4]string{"迪迦奥特曼", "赛罗奥特曼", "古加奥特曼"}
	var yourinput string
	fmt.Println("请输入要查找的光之子:")
	fmt.Scanln(&yourinput)
	// for i := 0; i < len(name); i++ {
	// 	if yourinput == name[i] {
	// 		fmt.Println("找到了，请相信光吧！")
	// 		break
	// 	} else if i == len(name)-1 {
	// 		//注意这种写法，不建议
	// 		fmt.Println("没有找到光之子")
	// 	}

	// }

	//第二种方式(推荐)
	var isContains = false
	for i := 0; i < len(name); i++ {
		if yourinput == name[i] {
			isContains = true

		}
	}

	if isContains {
		fmt.Println("找到了，请相信光吧！")
	} else {
		fmt.Println("没有找到光之子")
	}

	//3.二分法查找(有序+递归)
	var array2 = [5]int{12, 23, 45, 67, 89}
	//  (f(0) + f(n-1))/2 == ?
	twoF(array2, 0, 4, 23)
}

//冒泡算法
func bubbleSort(array *[6]int) {
	var temp int
	for j := 1; j < len(array); j++ {
		for i := 0; i < len(array)-j; i++ {
			//将大的元素往后移一位
			if (*array)[i] > (*array)[i+1] {
				//与array[i] > array[i+1] 不加*等效
				temp = (*array)[i+1]
				(*array)[i+1] = (*array)[i]
				(*array)[i] = temp
			}
		}
	}
}

//二分法查找算法,也可以用for循环来实现
// func twoF(array [5]int, start int, end int, findint int) {

// 	//如果左边游标大于或者等于右边游标时，就说明查找不到
// 	if start > end { //此处必须为 大于>，存在找到最后 start == end 【12,23】
// 		fmt.Println("查找不到哟")
// 		return
// 	}
// 	var mid = (end + start) / 2
// 	if findint > array[mid] {
// 		twoF(array, mid+1, end, findint)
// 	} else if findint < array[mid] {
// 		twoF(array, start, mid-1, findint)
// 	} else {
// 		//
// 		fmt.Println("找到了哟哈哈")
// 	}
// }

//二分查找的for实现,所有的递归都可以通过这种方式进行for转换，一般如果循环的次数是明确的考虑用for，不明确，用递归
//递归：空间换时间，for：时间换空间
func twoF(array [5]int, start int, end int, findint int) {

	for start <= end {
		mid := (start + end) / 2
		if array[mid] > findint {
			end = end - 1
		} else if array[mid] < findint {
			start = start + 1
		} else {
			fmt.Println("找到了")
			return
		}
	}
}

//排序和查找

// 排序是将一组数据，依指定的顺序进行排列的过程
// 排序的分类
// 1)内部排序:
// 指将需要处理的所有数据都加载到内部存储器中进行排序
// 包括(交换式排序、选择式排序法和插入式排序法)
// 2)外部排序法:
// 数据量过大，无法全部加载到内存中，需要借助外部存储进行排序。包括(合并排序法和直接合并排序法)
// 交换式排序属于内部排序法，是运用数据值比较后，依判断规则对数据位置进行交换，以达到排序的目的
// 2种方式:
// 1.冒泡排序法(Bubble sort)
// 2.快速排序法(Quick sort)

// 查找的分类
// 1)顺序查找：遍历判断相等即可
// 2)二分查找(Binaryfind)：必须是有顺序的

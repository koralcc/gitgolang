package main

import "fmt"

func main() {
	var arr1 [5]int
	//当定义完数组后,其实数组的各个元素都有了默认值(数据类型的默认值int0float0boolfalsestring"")
	fmt.Printf("数组的地址%p\n", &arr1)
	fmt.Printf("数组第一个元素的地址%p\n", &arr1[0]) //第一个元素的地址等于数组变量arr1的地址
	fmt.Printf("数组第二个元素的地址%p\n", &arr1[1]) //第二个元素 = 第一个元素地址+8(1个int占8个字节)

	//四种定义数组
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 = [3]int{1, 2, 3}
	var array3 = [...]int{1, 2, 3}
	var array4 = [3]string{1: "tome", 2: "jacke", 0: "marry"}
	fmt.Println("array1", array1)
	fmt.Println("array2", array2)
	fmt.Println("array3", array3)
	fmt.Println("array4", array4)
	//遍历
	for i, v := range array4 {
		fmt.Printf("第%d的值为%v", i, v)
	}

	//数组引用
	var arry5 = [5]int{1, 2, 3, 4, 5}
	test1(arry5)
	fmt.Println("值传递修改arry5后", arry5)

	test2(&arry5)
	fmt.Printf("arry5的地址为:%p", &arry5)
	fmt.Println("引用传递修改arry5后", arry5)

}

func test1(arry [5]int) {
	arry[1] = 100
	arry[2] = 200
}

func test2(arry *[5]int) {
	fmt.Printf("test2 arry地址%p\n", &arry)
	fmt.Printf("test2 arry的值%p,类型为%T\n", arry, arry)
	(*arry)[1] = 100
	(*arry)[2] = 200
}

//数组
// 数组可以存放多个同一类型数据.数组也是一种数据类型,在Go中,数组是值类型
// 1)数组的定义
// var 数组名 [数组大小]数据类型
// var a [5]int
// 赋初值a[0] =1
// 总结:(1)数组的地址可以通过数组名来获取 &arra,数组的地址就是首地址
// 	    (2)数组的第一个元素的地址,就是数组的首地址
// 	    (3)数组的各个元素的地址间隔是依据数组的类型决定的,比如int64->8,int32->4
//2)四种数组初始化方式
// var array [3]int = [3]int{1,2,3}
// var array  = [3]int{1,2,3}
// var array = [...]int{1,2,3}
// var array = [3]string{1:"tome",2:"jacke",0:"marry"}
// 3)数组的遍历
// (1)常规遍历
// (2)for-range遍历
// 4)数组的注意事项和细节
// (1)数组是多个相同类型数据的组合,一个数组一旦声明/定义了,其长度就是固定的,不能动态变化
// (2)var arr []int这时arr就是一个切片
// (3)数组中的元素可以是任何数据类型,包括值类型和引用类型,但是不能混用
// (4)数组创建后,如果没有赋值,有默认值
// 数值类型数组:默认为0
// 字符串数组:默认为""
// bool数组:默认为false
// (5)使用数组的步骤:1.声明数组并开辟空间2给数组各个元素赋值3使用数组
// (6)数组的下标是从0开始的
// (7)数组下标必须在指定范围内使用,否则报panic:数组越界,比如
// var arr[5]int则有效下标为0-4
// (8)Go的数组属于值类型,在默认情况下是值类型,因此会进行值拷贝
// (9)如果想在其他函数中,去修改原来的数组,可以使用引用传递(指针方式)
// (10)长度是数组的一部分,在传递函数参数时,需要考虑数组的长度,比如切片当数组传入,报错

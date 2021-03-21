package main

import "fmt"

func main() {
	//定义切片1
	var arry = [...]int{1, 2, 3, 4, 5}
	var sli = arry[1:3] //包含1不包含3
	fmt.Printf("arry地址为:%p\nsli地址为%p\narry[1]地址为%p\narry[2]地址为%p\nsli[0]地址为%p\nsli[1]地址为%p\n", &arry, &sli, &arry[1], &arry[2], &sli[0], &sli[1])
	//定义切片2
	var sli2 = make([]int, 2, 5)
	fmt.Println("切片sli2的数据是：", sli2) //[0,0],含有默认值
	//定义切片3
	var sli3 []int = []int{1, 3, 5}
	fmt.Println("切片sli3的数据为：", sli3)

	sli4 := arry[:]
	sli5 := arry[2:]
	sli6 := arry[:5]
	fmt.Println("sli4", sli4)
	fmt.Println("sli5", sli5)
	fmt.Println("sli6", sli6)

	sli7 := sli4[1:2]
	fmt.Println("sli7", sli7) //指向arry[1]的地址

	//slice扩容：append
	fmt.Printf("sli7[0]的地址%p\n", &sli7[0])
	sli8 := append(sli7, 3, 4) //新生成的一个底层数组，并将sli7重新指向新的这个底层数组
	fmt.Println("append后，sli7的值", sli7)
	sli7 = sli8
	fmt.Printf("append 后sli7[0]的地址%p\n", &sli7[0])
	//sli7 = append(sli7, 3, 4)
	fmt.Println("append后的sli7", sli7)
	sli7 = append(sli7, sli7...) //append后面也可以追加一个切片，...为将sli7转化为多参数
	//关于扩容的一个问题思考，在如上的扩容案例打印时，我们append前后的sli7即使在操作sli7 = sli8后，第一个地址还是与原来一致，
	//猜测有一个机制是，在扩容时，如果旧的数组没有引用指向了，还是在旧的内存上建新的数组进行扩容和重新赋值(如下结果就证实了这个猜想)
	var sliceA []int = []int{1, 2}
	fmt.Printf("append 前sliceA的地址为：%p\n", &sliceA[0])
	var sliceB = sliceA //如果没有这步，旧的底层数组被清除，新的数组复用这块地址进行新数组的创建
	fmt.Printf("sliceb地址%p,slicea地址%p", &sliceB[0], &sliceA[0])
	sliceA = append(sliceA, 3, 4)
	fmt.Printf("append 后sliceA的地址为：%p\n", &sliceA[0])
	fmt.Printf("append 后sliceB的地址为：%p\n", &sliceB[0])

	//切片的拷贝操作
	var sli10 []int = []int{1, 2, 3, 4, 5}
	var sli11 []int = make([]int, 1)
	fmt.Println(sli11)
	copy(sli11, sli10) //操作的切片多了少了多可以，少了就直接覆盖有的数量，copy是copy的值，不会有新的数组产生，与扩容的新建数组有差别
	fmt.Println("拷贝后的切片", sli11)

	//一个问题
	var a []int
	fmt.Printf("a的长度为%d,a的容量为%d,a的值为%v\n,a的地址为%p\n", len(a), cap(a), a, &a) //结果：0 0 []，不含默认值的空切片
	//a = []int{1, 2, 3}
	//fmt.Printf("a的长度为%d,a的容量为%d,a的值为%v\n,a的地址为%p\n", len(a), cap(a), a, &a) //结果：0 0 []，不含默认值的空切片
	//a[0] = 1 //panic: runtime error: index out of range [0] with length 0
	var b map[int]int
	fmt.Printf("b的长度为%d,b的容量为%d,b的值为%v\n,b的地址为%p\n", len(b), len(b), b, &b) //结果：0 0 []，不含默认值的空切片
	//b[1] = 1                                                                //panic: assignment to entry in nil map
	a = append(a, 1)
	fmt.Println(a)

	//斐波拉契数列，保存
	//1  1  2  3  5  8  f(n) = f(n-1) + f(n-2)//可以使用递归，这里也可以使用for循环
	var fblqN int = 10 // 记录保存多少个数
	fmt.Printf("%d个斐波拉契数为：%v", fblqN, getFblq(fblqN))

}

func getFblq(n int) []uint64 {
	var fblqSlice []uint64 = make([]uint64, n) //装斐波拉契数的切片
	fblqSlice[0] = 1
	fblqSlice[1] = 1

	for i := 2; i < n; i++ {
		fblqSlice[i] = fblqSlice[i-1] + fblqSlice[i-2] //从前往后推可以这样，从后往前的猴子问题依然可以这样做
	}
	return fblqSlice
}

//切片
// 1.切片的英文是slice
// 2.切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
// 3.切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len(slice)都一样
// 4.切片的长度是可以变化的，因此切片是一个可以动态变化的数组
// 5.切片定义的基本方法
// var 变量 []类型
// 比如：var a []int

//切片使用的三种方式
// 第一种方式：定义一个切片，然后切片去引用一个已经创建好的数组.[1,3]
// 第二种方式：通过make来创建切片
// 语法：var sli []int =make([]int,len,cap(len)),type是数据类型；len是大小;cap指定容量，可选
// 通过make方式创建切片，对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素
// 第三种方式:定义一个切片，直接就指定具体数组,使用原理类似make的方式
//  var slice []int = []int{1,3,5}

//方式1与方式2的区别(面试题)
// 方式1是直接引用数组，这个数组是事先存在的，程序员是可见的
// 方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的

//切片的遍历
// 切片的遍历和数组一样，也有两种方式
// 1)for循环常规方式遍历
// 2)for-range结构遍历

//切片的注意事项和细节
// 1)切片初始化时 var slice = arr[startindex:endindex]
// 说明：从arr数组下标为startindex，取到下标为endindex的元素(不包含arr[endindex])
// 2)切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
// var slice = arr[0:end]可以简写 var slice = arr[:end]
// var slice = arr[start:len(arr)]可以简写 var slice = arr[start:]
// var slice = arr[0:len(arr)]可以简写 var slice = arr[:]
// 3)cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
// 4)切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组或者make一个空间供切片来使用
// 5)切片可以继续切片 //新切片指向同一个底层数组
// 6)切片是引用类型，所以在方法传递时，遵守引用传递机制。传引用即可修改外面的变量

//关于切片特征的几个方法
// 1)append，扩容，往后追加元素，因为底层数组是一个值类型(不能修改，类似string)，所以每次在追加时都是新建一个底层数组，并将切片的引用指向新的数组，旧的数组被垃圾回收并释放
// 2)copy

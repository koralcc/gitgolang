package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//编写一个函数
	//随机猜数游戏
	//随机生成一个1--100的整数
	// 有十次机会
	// 如果第一次就猜中,提示"你真是个天才"
	// 如果2--3次猜中,提示"你很聪明,赶上我了"
	// 如果4--9次猜中,提示"一般般"
	// 如果最后一次猜中,提示"可算猜中啦"
	// 一次都每猜对提示"说你点啥好呢"
	//var randInt int64 //随机数
	var myInt int //我猜的数
	var count int //猜测的次数

	rand.Seed(time.Now().Unix()) //不给种子的话,默认每次执行都有个相同的默认值

	// var var randInt int64 //随机数
	// randInt = rand.Intn(100) + 1 //生成一个[0,100)的随机数 cannot use rand.Intn(100) + 1 (type int) as type int64 in assignmentgo
	randInt := rand.Intn(100) + 1 //生成一个[0,100)的随机数 randInt只要没定类型,右边都是可以默认转的 int和int64\int与float64等等
	fmt.Println("随机数为", randInt)
	//首先我们要知道，计算机不能产生绝对的随机数。只能产生伪随机数。伪就是有规律的意思。伪随机数就是计算机产生的随机数是有规律的。
	// 	那么计算机是怎么产生随机数的？
	// 当然是通过算法，这个算法是有映射关系的，如我放进1，他会出来一个特定的数
	// 统实现随机数是把当前的系统时间放进去，每次都不一样，所以可以实现。
	//label1:
	for {
		count++
		fmt.Println("请猜测一个数字:1")
		fmt.Scanln(&myInt)

		fmt.Printf("第%v次..", count)
		if count > 10 {
			return
		}

		if myInt != randInt {
			fmt.Println("没有猜对,继续...")
			continue
		}

		switch count {
		case 1:
			fmt.Println("你真是个天才")
			return //整个方法结束,如果在此处用break是退出switch,而不是退出forswitch默认带了break
			//break label(for)//也可以用break for的标签来退出for循环
		case 2, 3:
			fmt.Println("你很聪明,赶上我了")
			return
		case 4, 5, 6, 7, 8, 9:
			fmt.Println("一般般")
			return
		case 10:
			fmt.Println("可算猜中啦")
			return
		default:
			fmt.Println("说你点啥好呢")
			return
		}
	}

}

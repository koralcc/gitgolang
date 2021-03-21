package main

import (
	"fmt"
	"math"
)

func main() {

	//判断一个年份是否是闰年：1.能被4整除，但不能被100整除2.能被400整除
	var i int = 2020

	if (i%4 == 0 && i%100 != 0) || 1%100 == 0 {
		fmt.Printf("%v是一个闰年\n", i)
	} else {
		fmt.Printf("%v不是一个闰年\n", i)
	}

	//求ax xy2 + bx xy2 +c 方程的根，a,b,c分别为函数的参数，如果：b xy2 - 4ac >0 ,则有两个解；b xy2 - 4ac == 0，则有一个解；b xy2 - 4ac <0则无解
	var a, b, c float64 = 2, 5, 2
	if e := b*b - 4*a*c; e == 0 {
		r := (-b + math.Sqrt(e)) / 2 * a
		fmt.Printf("%v x2 + %vx + %v的根为%v", a, b, c, r)
	} else if e > 0 {
		r1 := (-b + math.Sqrt(e)) / 2 * a
		r2 := (-b - math.Sqrt(e)) / 2 * a
		fmt.Printf("%v x2 + %vx + %v的根为%v,%v", a, b, c, r1, r2)
	} else {
		fmt.Println("无解")
	}
}

//分支控制
// 分支控制就是让程序有选择的执行，有三种形式
// 1)单分支，if条件表达式{}  条件表达式是指能返回bool类型的常量、变量、有返回值的方法等
// 2)双分支，if条件表达式{}else{}
// 3)多分枝，if条件表达式{}else if条件表达式{}else{},else可以不写

// 双分支只会执行其中的一个语句
// 多分支要么只执行一个语句(有else)要么都不执行(没else)
// if条件语句中可以声明变量，也可以有括号，但官方不建议加括号

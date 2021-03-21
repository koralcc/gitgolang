package main

import "fmt"

func main() {
	// switch v := 3; {
	// case v == 3: //等于3或者等于4
	// 	fmt.Println(3, 4)
	// 	fallthrough
	// case v == 5, v == 6:
	// 	fmt.Println(5, 6)
	// default:
	// 	fmt.Println("no suit case")
	// }

	//1.使用switch把小写类型的char型转为大写(键盘输入)，只转换abcde，其他输出other
	var input byte
	fmt.Println("请输入一个字母:")
	//fmt.Scanln(&input) //这里的input存的都是实际值，故得用格式化scanf输入
	fmt.Scanf("%c\n", &input) //%n匹配输入的回车，回车会留在缓存中，下一个scan出现时，会自动补充这个回车，导致失效不然会影响到后面的2例子的scanln

	fmt.Println(input)
	switch input {
	case 'a', 'b', 'c', 'd', 'e':
		fmt.Printf("对应的大写字母为%c\n", 'a'-32)
	default:
		fmt.Println("other")
	}

	//2.对学生成绩大于60分得，输出合格，低于60分得输出不及格
	var score int
	fmt.Println("请输入一个数字:")
	fmt.Scanln(&score)
	fmt.Println("score", score)
	switch {
	case score > 60:
		fmt.Println("合格")
	case score < 60:
		fmt.Println("不合格")
	default:
		fmt.Println("数据不合理")
	}
}

//switch分支结构，细节

//1)case后是一个表达式(即：常量值、变量、一个有返回值的函数等都可以)
//   2)case后的各个表达式的值的数据类型必须和switch的表达式数据类型一致
//   3)case后面可以带多个表达式，使用逗号间隔。比如case表达式1，case表达式2
//   4)case后面的表达式如果是常量值(字面值)，则要求不能重复
//   5)case后面不需要带break，程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到则执行default
//   6)default语句不是必须的
//   7)switch后也可以不带表达式，类似多个if-else分支来使用
//   8)switch后也可以直接声明/定义一个变量，分号结束，不推荐
//   9)switch穿透fallthrough,如果在case语句块hh后增加fallthrough，则会继续执行下一个case，也叫switch穿透。

// switch和if的比较
// 1)如果判断的具体数量不多，而且符合整数、浮点数、字符、字符串这几种类型，建议用switch，简洁高效，例如 方向的上下左右，季节的春夏秋冬
// 2)其他情况，对区间的判断和结果为bool类型的判断，使用if，if的使用范围更广

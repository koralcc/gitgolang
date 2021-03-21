package main

import "fmt"

func main() {
	//main defer栈与sum的defer栈是独立的
	defer fmt.Println("main defer1") //5
	defer fmt.Println("main defer2") //4
	fmt.Println("main res", sum(1, 2))

}

func sum(a int, b int) int {
	var res int
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	//当函数执行完毕后(return后，defer在return后执行)，再从defer栈，按照先入后出的方式出栈
	defer fmt.Println("a=", a) //4  a = 1 连同数值一起压栈，即使后续做了自加操作，结果仍然是1
	defer fmt.Println("b=", b) //3  b = 2 连同数值一起压栈，即使后续做了自加操作，结果仍然是2
	defer func(res *int) {     //2
		*res += 100
		fmt.Println("defer res", *res) //res = 102
	}(&res)
	res = a + b
	a++
	b++
	fmt.Println("res=", res) //1
	return res
	//执行defer栈
}

//defer
//主要价值在，当函数执行完毕后，可以及时的释放函数创建的资源
//1)当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中【暂时称该栈为defer栈】，然后继续执行函数下一个语句
//2)当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行（遵守先入后出的机制）
//3)在defer将语句放入到栈时，也会将相关的值拷贝同时入栈

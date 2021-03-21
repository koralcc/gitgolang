package main

import "fmt"

func main() {

	//两种情况
	//1.值类型
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1的地址为：%p\n", &r1)
	//结果：r1的地址为：0xc000012440
	fmt.Printf("r1.leftUp地址为：%p,r1.leftUp地址为：%p\n", &r1.leftUp, &r1.rightDown)
	//结果：r1.leftUp地址为：0xc000012440,r1.leftUp地址为：0xc000012450
	fmt.Printf("r1.leftUp.x地址为：%p,r1.leftUp.y地址为：%p,r1.rightDown.y地址为：%p,r1.rightDown.y地址为：%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	//结果：r1.leftUp.x地址为：0xc000012440,r1.leftUp.y地址为：0xc000012448,r1.rightDown.y地址为：0xc000012450,r1.rightDown.y地址为：0xc000012458

	//2.指针/引用类型
	//r2有两个 *Point类型，这两个*Point类型的本身地址也是连续的
	//但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 12}, &Point{3, 5}, "aoteman"}
	fmt.Printf("r2本身地址为%p", &r2)
	fmt.Printf("r2.leftUp本身地址为%p,r2.rightDoan本身地址为%p,r2.name本身的地址为:%p\n",
		&r2.leftUp, &r2.rightDown, &r2.name)
	//结果：r2.leftUp本身地址为0xc000004500,r2.rightDoan本身地址为0xc000004508,r2.name本身的地址为:0xc000004510
	fmt.Printf("r2.leftUp.x地址为：%p,r2.leftUp.y地址为：%p,r2.rightDown.x地址为：%p,r2.rightDown.y地址为：%p\n",
		&r2.leftUp.x, &r2.leftUp.y, &r2.rightDown.x, &r2.rightDown.y)
	//他们指向的地址不一定是连续的，这个要看系统在运行时，是怎么分配的
	fmt.Printf("r2.leftUp指向的地址为%p,r2.rightDoan指向的地址为%p\n",
		r2.leftUp, r2.rightDown)
	//结果：r2.leftUp指向的地址为0xc0000140e0,r2.rightDoan指向的地址为0xc0000140f0
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
	name              string
}

//结构体的注意事项和使用细节
// 1)结构体的所有字段在内存中都是连续的(配合内存图，清晰)

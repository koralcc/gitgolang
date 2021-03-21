package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var heros HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	fmt.Println("排序前heros:")
	for _, v := range heros {
		fmt.Println(v)
	}
	//调用排序方法
	sort.Sort(heros)
	fmt.Println("排序后heros:")
	for _, v := range heros {
		fmt.Println(v)
	}

}

type HeroSlice []Hero

type Hero struct {
	Name string
	Age  int
}

//Len : 返回集合的长度
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less : 集合的排序方式：升序/降序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

// Swap : 集合的交换规则
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i] //等于上三句

}

//接口的最佳实践案例
//用sort.Sort()方法实现结构体切片排序

//接口 VS 继承
// 1)接口和继承解决的问题不同
// 继承的价值主要在于：解决代码的复用性和可维护性
// 接口的价值主要在于：设计，设计好各种规范(方法)，让其他自定义类型去试先这些方法
// 2)接口比继承更加灵活
// 接口比继承更加灵活，继承是满足 is- a的关系，而接口只需满足 like - a的关系，而接口只需满足
// 3)接口在一定程度上实现代码的解耦

//接口即体现了Golang的"""多态""",案例参照第一个usb插口

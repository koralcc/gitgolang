package main

import (
	"fmt"
	"strconv"
)

var sayCh = make(chan string, 300)
var endCh = make(chan struct{}, 1)

var intCh = make(chan int, 100)
var exitCh = make(chan bool, 1)

func main() {
	go WriteData(intCh)
	go ReadData(intCh)

	//for {
	select {
	case <-exitCh:
		return
	}
	//}

	// intCh <- 122
	// fmt.Println(<-intCh)

	// for {
	// 	v, ok := <-exitCh
	// 	if !ok {
	// 		return
	// 	}
	// 	fmt.Println(v)
	// }

	// iCh := make(chan bool, 1)
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Println("1")
	// 	}
	// 	iCh <- true
	// }()
	// <-iCh
}

func WriteData(intCh chan<- int) {
	for i := 0; i < 100; i++ {
		intCh <- i
	}
	close(intCh)
}

func ReadData(intCh <-chan int) {
	for {
		v, ok := <-intCh
		if !ok {
			//fmt.Println("ReadData", intD)
			break //切记这里不能用return，要用break
		}
		fmt.Println("ReadData", v)
	}

	// for v := range intCh {
	// 	fmt.Println("ReadData", v)
	// }//for range 与死循环读取，功能相似
	exitCh <- true
	//close(exitCh)//不关闭配合select使用
}

func catSay() {
	for i := 0; i < 100; i++ {
		sayCh <- "cat" + strconv.Itoa(i)
	}

	close(sayCh)
	endCh <- struct{}{}
	close(endCh)

}

func dogSay() {
	fmt.Println("dog")
}

func fishSay() {
	fmt.Println("fish")
	endCh <- struct{}{}
}

//有三个函数，分别打印"cat","dog","fish"
//要求每个函数都起一个goroutine，请按照"cat","dog","fish"的顺序打印，100次

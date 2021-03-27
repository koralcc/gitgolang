package main

import (
	"fmt"
)

func main() {

	//2.典例
	// 完成一个goroutine和channel协同工作的案例，具体要求
	// 1)开启一个writeData协程，向管道intChan中写入50个整数
	// 2)开启一个readData协程，从管道intChan中读取writeData写入的数据
	// 3)注意：writeData和readData操作的是同一管道
	// 4)主线程需要等待writeData和readData协程都完成工作才能退出【管道】
	intChan3 := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan3)
	go readData(intChan3, exitChan)

	// for v := range exitChan {
	// 	fmt.Println("读取到结束信号，结束！", v)
	// 	return
	// }
	for {
		v, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println("读取到结束信号，结束!", v)
	}

}

func writeData(intChan chan<- int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("channel写入：", i)
	}
	//写完关闭
	close(intChan)
}

func readData(intChan <-chan int, exitChan chan<- bool) {
	for v := range intChan {
		fmt.Println("channel读出：", v)
	}
	// for{
	// 	v,ok :=<- intChan
	// 	if !ok{
	// 		break
	// 	}
	// 	fmt.Println("channel读出：", v)
	// }
	exitChan <- true
	close(exitChan) //传递结束信号
}s

// 管道的关闭
// 管道的关闭是一个很重要的知识点，在阻塞中经常用到
// 使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据
// channel的遍历，channel支持for-range遍历，有两个细节需要注意
// 1)在遍历时，如果channel没有关闭，则会出现deadlock错误**
// 2)在遍历时，如果channel已经关闭，则会正确遍历数据，遍历完后，就会退出遍历**，读取到结束信号就退出循环

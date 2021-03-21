package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	go sayHello()
	time.Sleep(time.Second)
}

func test() {
	//这里我们可以使用defer+recover进行错误处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var testMap map[int]string
	testMap[0] = "hell0"
}
func sayHello() {
	for i := 0; i < 1000; i++ {
		fmt.Println("sayOk", i)
	}
}

//goroutine错误处理
//goroutine中使用recover，解决协程中出现panic，导致程序崩溃(如果不进行错误处理，任何一个协程出现错误，会导致所有的协程都进行不下去，包括，主线程)
// 如果我们起了一个协程，但是这个协程出现了panic，如果我们没有捕获这个panic，就会造成整个程序的崩溃，这时我们可以在goroutine中使用recover来捕获panic
// 进行处理，这样即使这个协程有问题，但是主线程不受影响，可以继续执行

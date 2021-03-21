package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
		stringChan <- fmt.Sprintf("str%d", i)
	}
	//传统的方式在遍历管道时，如果不关闭会阻塞而导致deadlock
	// for v := range intChan {
	// 	fmt.Println(v)
	// }//直接读取未关闭的channel会导致死锁

	for {
		select {
		case v := <-intChan: //注意：这里如果intChan一直没有关闭，它不会导致死锁，会自动到下一个case匹配
			fmt.Println("从intChan读取了数据", v)
		case v := <-stringChan:
			fmt.Println("从stringChan读取了数据", v)
		default:
			//channel都读到末尾
			fmt.Println("都娶不到了，程序员自定义的逻辑")
			return
		}
	}
}

//channel细节
// 1)channel可以声明为只读；或者只写性质
// 2)channel只读只写的案例
// 3)使用select可以解决从管道读取数据的阻塞问题，for-range遍历管道，必须要在写入的源头关闭管道，否则会报死锁
// 这样会对程序员产生一些困扰，什么时候关闭管道？select路由则不需要进行管道的关闭也可以读取到管道数据，不用
// 考虑管道的阻塞带来的问题
// 4)goroutine中使用recover，解决协程中出现panic，导致程序崩溃(如果不进行错误处理，任何一个协程出现错误，会导致所有的协程都进行不下去，包括，主线程)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var myMap = make(map[int]int, 10)
var lock sync.Mutex //互斥锁

func main() {
	//获取当前系统cpu的数量
	num := runtime.NumCPU()
	fmt.Println("当前系统的逻辑核心数：", num)
	//设置运行的cpu数目
	//runtime.GOMAXPROCS()
	// 1)go1.8后，默认让程序运行在多个栈上，可以不用设置
	// 2)go1.8前，还是要设置一下，可以更高效的利用cpu

	// 思考
	// 1.编写一个函数，来计算各个数的阶乘，并放入map中
	// 2.我们启动多个协程，统计的将结果放入到map中
	// 3.map应该做出一个全局的
	for i := 1; i <= 100; i++ {
		go getFactorial(i)
		//多个协程同时向map中写数据，会报错
		// fatal error: concurrent map writes
		// fatal error: concurrent map writes
		// 引出两个问题
		// 1.如何保证不冲突(加全局锁,与map声明处)
		// 2.如何确定goroutine的执行完毕时间
	}
	time.Sleep(time.Second * 5)
	fmt.Println("myMap", myMap)
}
func getFactorial(a int) {
	dstA := a
	for i := dstA; i >= 1; i-- {
		dstA *= i
	}
	lock.Lock()
	myMap[a] = dstA
	lock.Unlock()
}

//进程和线程说明
// 1)进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位
// 2)线程是进程的一个执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的基本单位。
// 3)一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行
// 4)一个程序至少有一个进程，一个进程至少有一个线程

//并行和并发
// 1)多线程程序在单核上运行，就是并发
// 2)多线程程序在多核上运行，就是并行

//Go协程和Go主线程
// 1)Go主线程（有程序员直接称为线程/也可以理解成进程）：一个Go线程上，可以起多个协程，你可以这样理解
// ，协程是轻量级的线程[编译器做优化]
// 2)Go协程的特点
// *有独立的栈空间
// *共享程序堆空间
// *调度由用户控制
// *协程是轻量级的线程

//goroutine-快速入门小结
// 1)主线程是一个物理线程，直接作用在cpu上，是重量级的，非常耗费cpu资源
// 2)协程从主线程开启的，是轻量级的线程，是逻辑态，对资源消耗相对小
// 3)Golang的协程机制是重要的特点，可以轻松的开启上万个协程，其他编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就凸显Golang在并发上的优势了

//MPG
// M:操作系统的主线程(是物理线程)
// P:写成执行需要的上下文(需要的环境、资源和运行时状态)
// G:协程

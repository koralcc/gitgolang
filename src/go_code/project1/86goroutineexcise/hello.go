package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//1.传统方法
	st := time.Now()
	getSuShu01(72)
	elapsed := time.Since(st)
	fmt.Println("传统函数执行完成耗时：", elapsed)
	//2.运用goroutine后
	fmt.Println(runtime.NumCPU(), "核cpu")
	st2 := time.Now()
	getSuShu02(72)
	elapsed2 := time.Since(st2)
	fmt.Println("传统函数执行完成耗时：", elapsed2)
}

func getSuShu01(dir int) {
	count := 0
	for i := 2; i < dir; i++ {
		if judgyIsPrime(i) {
			fmt.Println("素数：", i)
			count++
		}
	}
	fmt.Println("素数总数为：", count)
}

//创建8个goroutine进行统计
func getSuShu02(dir int) {
	intChan := make(chan int, dir)
	primeChan := make(chan int, dir)
	exitChan := make(chan bool, 8)
	//向channel中写入数据
	go func() {
		for i := 0; i < dir; i++ {
			intChan <- i
		}
		//写完以后关闭channel
		close(intChan)
	}()
	//开启8个goroutine进行数据的读取和素数判断
	for i := 0; i < 8; i++ {
		go func() {
			for {
				v, ok := <-intChan //只有关闭的channel这样去读，才会读到ok=false，否则会报死锁
				if ok && judgyIsPrime(v) {
					primeChan <- v
				}
				//读到关闭标识，说明当前goroutine读取结束,写入结束标识
				if !ok {
					exitChan <- true
					return
				}
			}
		}()
	}
	//单独开启一个goroutine关闭退出标识，读取所有8个goroutine结束标识
	go func() {
		//读取八次，相当于已经读完
		for i := 0; i < 8; i++ {
			<-exitChan //此处是会产生堵塞的
		}
		close(primeChan)
	}()

	//协程关闭处
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		//输出结果
		//fmt.Println("总数为:", len(primeChan))
		fmt.Println("素数为", v)
	}

}

//判断是否是素数
func judgyIsPrime(data int) bool {
	flag := true
	for j := 2; j < data; j++ {
		if data%j == 0 {
			flag = false
		}
	}
	return flag
}

//统计1-200000的数字中，哪些是素数

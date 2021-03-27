package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan struct{}, 2)
	t := time.NewTimer(time.Second * 1)
	go worker(ch)
	//var ch2 chan struct{}
	for {
		select {
		case <-ch:
			fmt.Println("读取到了哟")
		case <-t.C:
			ch = nil
		}
	}

}

func worker(ch chan struct{}) {
	ch <- struct{}{}
}

//1.向nil channel中，发/读，都会因堵塞而导致死锁
//2.没有close的channel循环获取，会导致死锁

//channel注意点
// 使用channel时有几个注意点：

// ·向一个nil channel发送消息，会一直阻塞；
// ·向一个已经关闭的channel发送消息，会引发运行时恐慌（panic）；
// ·channel关闭后不可以继续向channel发送消息，但可以继续从channel接收消息；
// ·当channel关闭并且缓冲区为空时，继续从从channel接收消息会得到一个对应类型的零值。

// 值得注意的是，在遍历时，如果channel 没有关闭，那么会一直等待下去，出现 deadlock 的错误；
// 如果在遍历时channel已经关闭，那么在遍历完数据后自动退出遍历。也就是说，for range 的遍历方式时阻塞型的遍历方式。

// select中有case代码块，用于channel发送或接收消息，任意一个case代码块准备好时，执行其对应内容；多个case代码块准备好时，随机选择一个case代码块并执行；所有case代码块都没有准备好，则等待；
// 还可以有一个default代码块，所有case代码块都没有准备好时执行default代码块。

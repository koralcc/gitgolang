package main

import "fmt"

func main() {
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	//2.看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	//3.向管道写入数据
	intChan <- 10
	intChan <- 90
	intChan <- 100
	//intChan <- 101 //fatal error: all goroutines are asleep - deadlock!
	//注意点，当我们给管道写入数据时，不能超过其容量3，否则会报死锁
	//4.看看管道的长度和容量(cap)
	fmt.Printf("channel len=%v,cap=%v\n", len(intChan), cap(intChan)) //len=3,cap=3
	//5.从管道中读取数据,取过数据后，管道的长度(len)会变小
	<-intChan
	fmt.Printf("channel len=%v,cap=%v\n", len(intChan), cap(intChan)) //len=2,cap=3
	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报deadlock
	num1 := <-intChan
	num2 := <-intChan
	//num3 := <-intChan//fatal error: all goroutines are asleep - deadlock!
	fmt.Println(num1, num2 /*, num3*/)

	//类型断言的一个案例
	var allChan chan interface{} = make(chan interface{}, 5)
	allChan <- 10
	allChan <- "tom"
	cat := Cat{"小黑猫", "黑色"}
	allChan <- cat
	//我们希望获取到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newCat := <-allChan //这里的newCat其实是interface{}的引用指向cat的实例
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	//fmt.Println(newCat.Name)//newCat.Name undefined (type interface {} is interface with no methods)
	a := newCat.(Cat)
	fmt.Println("猫名", a.Name)
	//管道的遍历
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}

	//如果遍历时，没有关闭管道会报死锁all goroutines are asleep - deadlock!
	close(intChan2) //不关闭的话，for会一直等待在这里，系统判定为死锁
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	//1.管道的遍历
	intChan3 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan3 <- i
	}

	//如果遍历时，没有关闭管道会报死锁all goroutines are asleep - deadlock!
	close(intChan3) //不关闭的话，for会一直等待在这里，系统判定为死锁
	for v := range intChan3 {
		fmt.Println("v=", v)
	}
}

type Cat struct {
	Name  string
	Color string
}

//channel
// 1)channel本质是一个数据结构-队列
// 2)数据是先进先出【FIFO】
// 3)线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
// 4)channel有类型的，一个string的channel只能存放string类型数据

//声明
// var 变量 chan 数据类型
// var intChan chan int
// var mapChan chan map[int]string
// var perChan chan Person
// var perChan2 chan *Person
// //说明
// 1)channel是引用类型
// 2)channel必须初始化才能写入数据，即make后才能使用
// 3)管道是有类型的，intChan只能写入整数int

//注意事项
// 1)channel中只能存放指定的数据类型
// 2)channel的数据放满后，就不能再放入了
// 3)如果从channel取出数据后，可以继续放入
// 4)在没有使用协程的情况下，如果channel数据取完了，再取，就会报deadlock

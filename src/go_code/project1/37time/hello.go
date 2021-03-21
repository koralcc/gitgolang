package main

import (
	"fmt"
	"time"
)

func main() {
	//1.获取当前时间
	now := time.Now()
	fmt.Printf("nowtype=%T,now=%v\n", now, now)
	//2.获取年月日时分秒
	fmt.Println("Year", now.Year())
	fmt.Println("Month", now.Month())
	fmt.Println("Day", now.Day())
	fmt.Println("Houre", now.Hour())
	fmt.Println("Minute", now.Minute())
	fmt.Println("Second", now.Second())
	//3.格式化日期和时间
	//(1)第一种方式
	fmt.Printf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())         //按对应格式打印
	str := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()) //获取对应的格式
	fmt.Println(str)
	//(2)第二种方式
	fmt.Printf(now.Format("2006-01-02 15:04:05")) //固定的数字 2006-01-02 15:04:05
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Printf(now.Format("15:04:05"))
	fmt.Printf(now.Format("05"))
	fmt.Println()
	//4.time常量
	//对于需要休眠多少间隔后的操作,用time常量
	time.Sleep(time.Second * 2) //这里只能乘以整数不能乘以小数,因为golang int * float32会报错
	//5.获取时间戳
	fmt.Println("获取时间戳", now.Unix())
	fmt.Println("获取纳秒时间戳", now.UnixNano())
	//练习获取执行代码时间time.Now().Unix() - time.Now().Unix()
}

//time

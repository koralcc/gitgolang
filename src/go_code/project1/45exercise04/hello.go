package main

import (
	"fmt"
	"time"
)

//三天打鱼两天晒网,如果从1990-1-1开始,给定一个日子,这个日子是打鱼还是晒网  打鱼打鱼打鱼 晒网晒网
func main() {
	//用间隔天数对5取模,即能得到是在五天的第几天,1-3打鱼,4-5晒网
	//time 包有个函数 Parse 可以将时间字符串解析成 Time 对象，而 Time 对象有个 Sub 方法可以计算与某个时间的差，
	//返回值是 Duration 对象，而 Duration 有一个 Hours 方法，除以 24 就是天数了
	//获取1990-1-1的time对象

	var endtimestr string
	startTime, _ := time.Parse("2006-01-02", "1990-01-01")
	fmt.Println(startTime)

	fmt.Println("请输入一个日期,格式为XX-XX-XX:")
	fmt.Scanln(&endtimestr)
	endtime, _ := time.Parse("2006-01-02", endtimestr)

	//获取两个时间之间的差
	duration := endtime.Sub(startTime)

	days := duration.Hours() / 24 //即为两个时间隔得天数

	fmt.Println("相隔多少天", days)
	whicd := int64(days) % 5
	fmt.Println("是第几天", whicd)
	if whicd >= 1 && whicd <= 3 {
		fmt.Println("打鱼")
	} else {
		fmt.Println("晒网")
	}
}

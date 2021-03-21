package main

import "fmt"

//1.循环打印输入的月份日期.[使用continue实现]
//要有判断输入的月份是否错误的语句
// 31天:1,3,5,7,8,10,12
// 30天:4,6,9,11
// 28天:非闰年 2
// 29天:闰年 2月
func main() {

	var year, month, day int64

	for {

		fmt.Println("请输入年份:")
		fmt.Scanln(&year)
		if year < 0 {
			fmt.Println("年份不正确")
			continue
		}

		fmt.Println("请输入月:")
		fmt.Scanln(&month)
		if month <= 0 || month > 12 {
			fmt.Println("月份不正确")
			continue
		}

		fmt.Println("请输入日:")
		fmt.Scanln(&day)
		//日期是否准确的判断
		isLeap := judyIsLeap(year)
		var dayRight bool
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if day <= 31 && day >= 1 {
				dayRight = true
			}
		case 4, 6, 9, 11:
			if day <= 30 && day >= 1 {
				dayRight = true
			}
		case 2:
			//2月份
			if day >= 1 && (isLeap && day <= 29) || (!isLeap && day <= 28) {
				dayRight = true
			}
		}
		if dayRight == false {
			fmt.Println("日期不正确!!")
			continue
		}

		fmt.Printf("您输入的日期正确,为%v-%v-%v:\n", year, month, day)

		year = 0
		month = 0
		day = 0
	}

}

//根据年,判断是否为闰年
func judyIsLeap(year int64) bool {
	//根据年判断是否是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		//闰年 2月份有29天
		return true
	} //else {
	//2月份有28天
	return false
	//}
}

//阶段性大练习

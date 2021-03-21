package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 int8 = 88
	var f1 float32 = 12.111
	var c1 byte = 'A'
	var b1 bool = tru
	var s1 string //空的string
	s1 = fmt.Sprintf("%d", i1)
	fmt.Printf("s1type=%T,s1=%v\n", s1, s1)
	s1 = fmt.Sprintf("%f", f1)
	fmt.Printf("s1type=%T,s1=%v\n", s1, s1)
	s1 = fmt.Sprintf("%c", c1)
	fmt.Printf("s1type=%T,s1=%v\n", s1, s1)
	//%T:类型，%t布尔
	s1 = fmt.Sprintf("%t", b1)
	fmt.Printf("st1type=%T,s1=%q\n", s1, s1)

	var i2 int8 = 89
	var f2 float64 = 16.111
	// var c2 byte = 'B'
	var b2 bool = false
	var s2 string //空的string
	s2 = strconv.FormatInt(int64(i2), 10)
	fmt.Printf("s2type=%T,s2=%v\n", s2, s2)
	s2 = strconv.FormatFloat(f2, 'f', 10, 64)
	//'f'格式，10：表示小数点保留10位，64：表示这个数是float64
	fmt.Printf("s2type=%T,s2=%v\n", s2, s2)
	// s2 = strconv.FormatInt(int64(c2), 10)
	// fmt.Printf("s2type=%T,s2=%c\n", s2, s2)
	s2 = strconv.FormatBool(b2)
	fmt.Printf("s2type=%T,s2=%v\n", s2, s2)

}

//基本数据类型转string
// 方式一:fmt.Sprintf("%参数",表达式)【推荐这个，灵活】
// 1)参数需要和表达式的数据类型相匹配
// 2)fmt.Sprintf会返回转换后的字符串
// 方式二:使用strconv包函数
// strconv.FormatInt
// strconv.FormatFloat
// strconv.FormatBool
// strconv.Itoa

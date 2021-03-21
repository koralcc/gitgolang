package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//1.统计字符串的长度: len(str)
	var str = "hello中国"
	//Golang的编码统一为utf-8(ascii的字符(字母和数字)占一个字节,汉字占3个字节)
	fmt.Println("字符串长度:", len(str)) //统计的是字节数

	//2.字符串遍历,同时处理有中文的问题: r:= []rune(str)
	var str2 = "hello中国"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		//fmt.Println("字符=", str2[i]) //字符的unicode码,104 101等
		fmt.Printf("字符=%c\n", str3[i])
	}

	//3.字符串转整数: n,err := strconv.Atoi("123")
	var str4 string = "123"
	int4, err := strconv.Atoi(str4)
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("字符串转整数结果:%v\n", int4)
	}

	//4.整数转字符串: n:= strconv.Itoa(123)
	var int5 = 123
	str5 := strconv.Itoa(int5)
	fmt.Printf("整数转字符串结果:%v\n", str5)

	//5.字符串转byte切片:var bytesf = []byte("hello go")
	var str6 = "hello 中国"
	var bytes = []byte(str6)
	fmt.Printf("字符串转byte切片结果:%v", bytes)

	//6.byte切片转字符串:var str = string([]byte{123,124,124,234})
	str7 := string([]byte{97, 98, 99})
	fmt.Printf("byte切片转字符串结果:%v\n", str7)

	//7.10进制转2.8.16进制字符串:str:= strconv.FormatInt(123,2)
	var int9 = 100 //默认为int类型,虽然是64位但是int != int64,需要转化,将来可能还有int128,所以是两种类型(同样的一份代码,在32位和64位的机器上)
	fmt.Printf("100转2进制结果:%v\n", strconv.FormatInt(int64(int9), 2))
	fmt.Printf("100转8进制结果:%v\n", strconv.FormatInt(int64(int9), 8))
	fmt.Printf("100转16进制结果:%v\n", strconv.FormatInt(int64(int9), 16))

	//8.查找子串是否存在指定的字符串中:strings.Contains("seafood","foo")
	var str10 = "seafood"
	var str11 = "fo0"
	var contsChild = strings.Contains(str10, str11)
	fmt.Printf("父串%v含有子串%v:%v\n", str10, str11, contsChild)

	//9.统计一个字符串有几个指定的子串:strings.Count("cheese","e")
	var str12 = "cheese"
	var str13 = "e"
	var countStr = strings.Count(str12, str13)
	fmt.Printf("父串%v的还有%v个%v\n", str12, countStr, str13)

	//10.不区分大小写的字符串比较(==是区分字母大小写) :strings.EqulFold("abc","abc")
	var str14 = "abc"
	var str15 = "AbC"
	fmt.Printf("两个字符串不区分大小比较相同:%v", strings.EqualFold(str14, str15))
	fmt.Println("abc" == "Abc")

	//11.返回子串在字符串第一次出现的index值,如果没有返回-1:strings.Index("CC_nt","nt")
	str16 := "Cc_Nt"
	str17 := "Nt"
	fmt.Printf("子串%v在字符串%v中第一次出现的位置%v\n", str17, str16, strings.Index(str16, str17))

	//12.返回子串在字符串中最后一次出现的index,如没有返回-1:strings.LastIndex("cc_Nt","nt")
	str18 := "Cc_Nt"
	str19 := "Nt"
	fmt.Printf("子串%v在字符串%v中最后一次出现的位置%v\n", str19, str18, strings.LastIndex(str18, str19))

	//13.将指定的子串替换成另外的子串:strings.Replace("hellolaolbbbol", "ol", "xx", n),n为替换第几个,-1为全部替换
	fmt.Println("将ol替换成xx", strings.Replace("hellolaolbbbol", "ol", "xx", -1))

	//14.按照给定的某个字符,为分割标识,将一个字符串拆分成字符串数组,strings.Split("hello,word,ok",",")
	fmt.Println("将 hello,word,ok 按,拆分成数组", strings.Split("hello,word,ok", ","))

	//15.将字符串的字母进行大小写的转换:strings.toLower("Go")
	//var str20 = "go"//0xc000088100
	// str20 = "tt"/0xc000088100//此处的修改其实相当于new了一个str20新的空间,只是空间地址还是刚才这个
	fmt.Println("将Go转化成小写", strings.ToLower("Go"))
	fmt.Println("将Go转化成小写", strings.ToUpper("Go"))

	//16.将字符串左右两边的空格去掉:strings.TrimSpace("  hello    ")
	fmt.Println("去除字符串  Hello  左右两边的空格", strings.TrimSpace(" hello   hello  "))

	//17.将字符串左右两边指定的字符去掉 strings.Trim(":hello:",":")
	fmt.Println("将字符串左右两边的:去掉", strings.Trim(":hello:", ":"))

	//18.将字符串左边指定的字符去掉 strings.TrimLeft(":hello:",":")
	fmt.Println("将字符串左边的:去掉", strings.TrimLeft(":hello:", ":"))

	//19.将字符串左边指定的字符去掉 strings.TrimLeft(":hello:",":")
	fmt.Println("将字符串右边的:去掉", strings.TrimRight(":hello:", ":"))

	//20.判断字符串是否以指定的字符串开头:strings.Hasprefix("ftp://168.8.8.1","ftp")
	fmt.Println("判断字符串ftp://168.8.8.1是否以ftp开头", strings.HasPrefix("ftp://168.8.8.1", "ftp"))

	//21.判断字符串是否以指定的字符串结尾:strings.HasSuffix("www.baidu.com","com")
	fmt.Println("判断字符串www.baidu.com是否以com结尾", strings.HasSuffix("www.baidu.com", "com"))

}

//字符串中常用的系统函数

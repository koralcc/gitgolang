package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "123455" //分配到只读内存段
	fmt.Println("[]byte str:", []byte(str))
	//var sli []byte
	//获取字符串底层数组的指针
	//strPtr := unsafe.Pointer(&(sli[0]))
	//获取str中 Data指针
	//p := (*reflect.StringHeader)(unsafe.Pointer(&str))
	//获取字符串指向的底层数组地址
	// strptr := unsafe.Pointer(p.Data)
	// fmt.Printf("p:%p\n", strptr) //str指向的底层数组地址
	//fmt.Println("p Len", *(*string)(strptr))
	//fmt.Printf("&str[0]:%p", &(str[0]))         //&(str[0] not allowed
	//将string底层数组指针直接赋给切片
	// sli = *(*[]byte)(strptr)
	// fmt.Println("sli", sli)
	b := StringToSlice(str)
	//b[0] = 10 // unexpected fault address 0x4cd587,只读内存段不可写
	fmt.Println(b)
	b2 := []byte(str)
	b2[0] = 10
	fmt.Println(b2)
	//b[0] = 10
}

func StringToSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data //直接将string底层数组指针指向byte
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

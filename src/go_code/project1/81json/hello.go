package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//序列化结构体
	st := student{
		"tome",
		true,
		18,
		99,
	}
	str, err := json.Marshal(st)
	if err != nil {
		fmt.Println("序列化报错", err)
	}
	fmt.Println("st的json串", string(str))
	//map进行序列化
	m1 := make(map[string]interface{})
	m1["name"] = "牛魔王"
	m1["age"] = 600
	m1["skill"] = "牛魔拳"
	str1, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("序列化报错", err)
	}
	fmt.Println("m1的json串", string(str1)) //从结果可以看出map的key是没有顺序的

	//反序列化结构体
	webStr := "{\"name\":\"tome\",\"Sex\":true,\"Age\":18}"
	var stUnm student
	//json.Unmarshal([]byte(webStr), stUnm) //这里必须传引用，不然改变不了
	json.Unmarshal([]byte(webStr), &stUnm)
	fmt.Println("反序列化后的stu", stUnm)
	//反序列化map，map的反序列化不需要make，被封装到了unmashal里面
	webMapStr := "{\"age\":600,\"name\":\"牛魔王\",\"skill\":\"牛魔拳\"}" //程序自己读到的是不需要加\的
	var m1Unm map[string]interface{}
	json.Unmarshal([]byte(webMapStr), &m1Unm)
	fmt.Println("m1的反序列化后", m1Unm)
}

type student struct {
	Name  string `json:"name"` //注意tag的使用格式是 json:"name" ,实现的原理是反射
	Sex   bool
	Age   float64
	score float64 //不能将属性定义为私有小写，序列化相当于在其他包进行私有属性的调用，读取不到
}

//序列化与反序列化
//1)序列化 json.Marshal()
//2)反序列化 json.Unmarshal()

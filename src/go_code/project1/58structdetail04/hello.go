package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//1.创建一个monster变量
	monster := Monster{"白骨精", 190, "变身"}
	//2.将monster变量序列化为json格式字符串
	str, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化错误", err)
	}
	fmt.Println("jsonStr", string(str)) //str:[]byte
	//要是把Monster的字段定义成小写，那么Json.Marshal外面的包将引用不到这个字段,
	//所以这里也得到一个规律，需要序列化的结构体，必须定义成大写

}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//结构体的注意事项和使用细节(详情见下面几节)
// 	)struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化

package main

import (
	"fmt"
	"sort"
)

func main() {

	var m map[string]string                     //map零值(不赋值的默认值):nil
	var s []int                                 //slice零值(不赋值的默认值):nil
	fmt.Printf("m的地址%p,s的地址%p,%v\n", &m, &s, m) //这两种底层为对象的结构看成是整个对象的地址
	//m["1"] = "1"                              //assignment to entry in nil map
	m = make(map[string]string)
	fmt.Printf("m的地址%p,s的地址%p,%v", &m, &s, m)

	//存三个学生的数据，每个学生有age和name
	studentMap := make(map[string]map[string]string)
	studentMap["stuo1"] = make(map[string]string, 2)
	studentMap["stuo1"]["name"] = "tom"
	studentMap["stuo1"]["age"] = "18"

	studentMap["stuo2"] = make(map[string]string, 2) //这里必须初始化
	studentMap["stuo2"]["name"] = "jimmy"
	studentMap["stuo2"]["age"] = "78"

	fmt.Println(studentMap)

	m2 := make(map[string]string, 3)
	fmt.Println("m2长度为", len(m2)) //0
	m2["1"] = "北京"
	m2["2"] = "上海"
	m2["3"] = "广州"
	m2["4"] = "深圳"
	fmt.Println("修改前", m2)
	m2["4"] = "杭州"
	fmt.Println("修改后", m2)
	delete(m2, "4")
	fmt.Println("删除后", m2)
	fmt.Println("m2长度为", len(m2)) //3,有值的长度,key--value对的数目
	v, ok := m2["4"]
	if ok {
		fmt.Println("对应的值为", v)
	} else {
		fmt.Println("不存在对应的值", v == "")
	}

	//map切片,map的个数可以动态变化
	var monsterArry = make([]map[string]string, 2) //切片make时，长度是必填项
	monsterArry[0] = make(map[string]string)
	monsterArry[0]["name"] = "白骨精"
	monsterArry[0]["age"] = "18"
	monsterArry[1] = make(map[string]string)
	monsterArry[1]["name"] = "蜘蛛精"
	monsterArry[1]["age"] = "18"
	fmt.Println("monsterarry", monsterArry)
	//这种情况增加第三个时， 因为外层切片定义的是，[2]会报越界
	monsterNew := map[string]string{
		"name": "红孩儿",
		"age":  "122",
	}
	monsterArry = append(monsterArry, monsterNew)
	fmt.Println("添加后的monster", monsterArry)

	//思考一个问题，切片长度声明为2，直接用下标2添加第三个元素，会不会通过？
	// var s1 = make([]int, 2)
	// s1[0] = 1
	// s1[1] = 2
	// s1[2] = 3
	// fmt.Println(s1) //index out of range [2] with length 2
	// //切片的扩容是通过append来实现的

	//map排序
	var m4 = map[int]string{
		2: "唐生",
		1: "孙悟空",
		3: "猪八戒",
	}
	var arrykeys []int = make([]int, 3)
	for i, _ := range m4 {
		arrykeys = append(arrykeys, i)
	}
	//排序
	sort.Ints(arrykeys) //此处就可以看出map是引用类型
	for _, v := range arrykeys {
		fmt.Println(v, m4[v])
	}
	s1 := student{
		name: "1",
		age:  87,
	}
	fmt.Println(s1)
	testStruct(s1)
	fmt.Println("modi", s1)
}

type student struct {
	name string
	age  float64
}

func testStruct(s student) {
	s.name = "xiuxxiuxu"
}

//map
// 1)map在使用前一定要make
// 2)map的key是不能重复的，如果重复了，则以最后这个key-value为准
// 3)map的value是可以相同的
// 4)map的key-value是无序的
// 5)make内置函数make，不带长度时，默认为1

//map的使用方式
// 三种
// 方式一
// var m map[string]string
// m = make(map[string]string,10)
// 方式二
// var m = make(map[string]string)
// 方式三
// var m = map[string]string{
// 	"1" : "1",
// 	"2" : "2",
// }

//map的增删改查操作
// map增加和更新
// map[key] = value //如果key还没有，就是增加，如果key存在就是修改
// map删除
// delete(map,"key"),delete 是一个内置函数，如果key存在，就删除key-value，如果key不存在，不操作，但是也不报错
// 如果我们要删除map的所有key，没有一个专门的方法一次删除，可以遍历一下key，逐个删除
// 或者map=make(),make一个新的，让原来的成为垃圾，被GC回收
// map查找
//  v,ok := m["key"]
//  ok,是否存在，存在true，不存在false
//  v存在对应的值

//map的遍历 只有一种方式 for---range

//map切片
//切片的数据类型如果是map，则我们称为slice of map，map切片，这样使用则map的个数就是可以动态改变了

//map排序
// 1)Golang中没有一个专门的方法针对map的key进行排序
// 2)Golang中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历得到的暑促结果可能都不一样
// 3)Golang中map的顺序，是先将key进行排序，然后根据key值遍历输出即可

// map的使用细节
// 1)map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
// 2)map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长键值对
// 3)map的value也经常使用struct类型，更适合管理复杂的数据（比前面练习中value是一个map的更好）

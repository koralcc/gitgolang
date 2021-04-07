package main

func main() {
	var ints = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//准备两个游标
	front := 1
	end := 2
	//记录front--end中间的砖块数量
	var countQut = 0
	//记录front--end中间间的总体积)
	var countAll = 0
	for {

		if ints[end] >= ints[front] {
			front = end //进行下一次循环
		}
		end++
	}
}

//面试题 17.21. 直方图的水量

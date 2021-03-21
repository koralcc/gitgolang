package main

import "fmt"

func main() {
	//十进制 decimal
	var d int = 12
	//二进制 Binary
	fmt.Printf("d的二进制为%b\n", d)
	//八进制 octonary
	var o int = 014
	fmt.Printf("o的八进制值为%v\n", o)
	//16进制 hexadecimal
	var h int = 0x00A
	fmt.Printf("h的十六进制值为%v\n", h)
}

//进制
// 对于整数，有四种表示方式
// 1)二进制：0,1，满2进1
//   在golang中，不能直接使用二进制来表示一个整数，他沿用了c的特点
// 2)十进制：0-9，满10进1
// 3)八进制：0-7，满8进1，以数字0开头
// 4)十六进制：0-9及A-F，满16进1，以0X或0x开头表示
// 此处的A-F不区分大小写

//	十进制		  十六进制		八进制		二进制
//		0		    	0	        0			0
//		1				1			1			1
//		2				2			2		    10
//		3				3			3		    11
//		4				4			4		    100
//		5				5			5			101
//		6				6			6			110
//		7				7			7			111
//		8				8			10			1000
//		9				9			11			1001
//		10				A			12			1010
//		11				B 			13			1011
//		12				C			14			1100
//		13				D			15			1101
//		14				E			16			1110
//		15				F			17			1111
//		16				10			20			10000

//进制的相互转化
// 1)其他进制转十进制
// 规则：从最低位开始(右边的)，将每个位上的数据提取出来，乘以2/8/16的(位数-1)次方，然后求和
// 2)十进制转其他进制
// 规则：将该数不断除以2/8/16，直到商位0为止，然后将每步得到的余数倒过来，就是对应的进制数
// 3)二进制转其他进制
//   (1)二进制转八进制
// 	规则：将二进制数每三位一组(从低位开始组合),转成对应的八进制数即可。(对照上图)
//   (2)二进制转16进制
// 	规则：同上将二进制数每四位一组,转成对应的十六进制数即可
// 4)其他进制转二进制
//   (1)八进制转二进制
// 	规则：将八进制数每一位，转成对应的一个3位的二进制数即可
//   (2)十六进制转二进制
//     规则：将十六进制数每一位，转成对应的一个4位的二进制数即可
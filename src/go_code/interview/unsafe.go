package main
import(
	"fmt"
	"unsafe"
)
func main(){

	var a int = 10
	var b int = 20
	fmt.Println("a:",a)
	fmt.Println("b:",b)

	fmt.Printf("&a type:%T\n",(*int)(unsafe.Pointer(&a)))
	//c := &a + &b //连个地址之间是不能运算的

	stu:= &Student{}
	//获取stu的指针值,用stu指针+b的偏移量得到b的地址，根据地址给b赋值
	var bint uintptr = uintptr(unsafe.Pointer(stu)) + unsafe.Offsetof(stu.b)
	fmt.Println("bintsize:",unsafe.Sizeof(bint))
	//将b的地址整数转化为指针类型
	var bptr unsafe.Pointer = unsafe.Pointer(bint)
	fmt.Println("*bptr",bptr)
	//将unsafe.Pointer转化为b的 *int类型
	bptrint := (*int)(bptr)
	//将b赋值
	*bptrint = 20000
	fmt.Println(stu.b)

	//---------------内存对齐-------
	stu2 := Student{10,20,30,30}
	fmt.Println("stu.size:",unsafe.Sizeof(stu2))

	//---------------内存对齐-------
	fmt.Println("stu的内存对齐方针为：",unsafe.Alignof(stu2))
	fmt.Printf("&a:%p\n&b:%p\n&c:%p\n&d:%p\n",&stu.b,&stu.c,&stu.d,&stu.a)

	// fmt.Printf("stu:%p\n",stu)
	// fmt.Printf("stu.b:%p\n",&(stu.b))
	// fmt.Println(unsafe.Offsetof(stu.b))
	//fmt.Println("stu:",stu)

}

type Student struct{
	a  int8 
	b  int64
	c  int32  
	d  int16  
	
	// b  int64
	// c  int32  
	// d  int16  	
	// a  int8  //按这个顺序，student的长度变成了16，更省内存 
}

//unsafe包
//uintptr
// · 一个足够大的无符号整型， 用来表示任意地址。
// · 可以进行数值计算。

//·  unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
//·  而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
//·  unsafe.Pointer 可以和 普通指针(*int *string *bool) 进行相互转换；
//·  unsafe.Pointer 可以和 uintptr 进行相互转换。



//2.内存对齐
//cpu从内存读取某个数据，是按一组机器字来读的(32位，每次读取4个字节；64位，每次读取8个字节)
//对于一个结构体，而言，如果变量在内存的分布不进行内存对齐，会增加cpu对内存的io次数，这个时间相对于cpu来说是很长的，所以需要进行内存对齐

//内存对齐的规则：2个
// 1.变量相对地址是对齐边界的整数倍，即：32位需要%4=0,64位需要%8=0
// 2.结构体总长度是最大对齐边界的整数倍，即：32位按4，8，12，16扩展长度；64位按照8，16，24，32扩展长度
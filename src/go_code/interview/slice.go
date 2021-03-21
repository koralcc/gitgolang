package main
import(
	"fmt"
	"unsafe"
)
func main() {
	//1.用数组构造的切片，与数组共享内存空间
	var arr [10]int
	var sli = arr[5:6]
	fmt.Println("&arr[5]:",&arr[5])
	fmt.Println("&sli[5]:",&sli[0])
	fmt.Printf("sli len:%d,cap:%d\n",len(sli),cap(sli))

	//2.切片的可变长通过append方法实现,append函数执行时，会判断切片容量是否能够
	//存放新增元素，如果不能，则会重新申请存储空间。
	var sli2 = make([]int,3,5)
	//var sli2 []int
	fmt.Printf("&slice2:%p,cap:%d,len:%d\n",&sli2,cap(sli2),len(sli2))
	newSlice4 := append(sli2,11,22)
	fmt.Printf("&slice4:%p,cap:%d,len:%d\n",&newSlice4,cap(newSlice4),len(newSlice4))
	newSlice5 := append(newSlice4,4)
	fmt.Printf("&slice5:%p,cap:%d,len:%d\n",&newSlice5,cap(newSlice5),len(newSlice5))

	var slice []int
    slice = append(slice, 1, 2, 3,4)
	fmt.Printf("&slice:%p,cap:%d,len:%d\n",slice,cap(slice),len(slice))
	fmt.Println(&slice[0])
    newSlice := AddElement(slice, 4)	
	fmt.Printf("&newSlice:%p,cap:%d,len:%d\n",&newSlice,cap(newSlice),len(newSlice))
	fmt.Println(&newSlice[0])
    fmt.Println(&slice[0] == &newSlice[0])

	stu := student{}
	fmt.Println("stu.size:",unsafe.Sizeof(stu))

	//---------------内存对齐-------
	fmt.Println("stu的内存对齐方针为：",unsafe.Alignof(stu))
}

func AddElement(slice []int, e int) []int {
    return append(slice, e)
}

type student struct{
	a  int8  
	b  int16 
	c  int64  
	d  int16    
	 
	//0 0 0 0 0
	 
}


//1.切片的扩容通过append
//2.append函数执行时，会判断切片容量是否能够存放新增元素，如果不能，则会重新申请存储空间。容量小于1024每次扩容乘以2，容量大于1024每次扩容乘以1.25

// · 每个切片都指向一个底层数组
// · 每个切片都保存了当前切片的长度、底层数组可用容量
// · 使用len()计算切片长度时间复杂度为O(1)，不需要遍历切片
// · 使用cap()计算切片容量时间复杂度为O(1)，不需要遍历切片
// · 通过函数传递切片时，不会拷贝整个切片，因为切片本身只是个结构体而已
// · 使用append()向切片追加元素时有可能触发扩容，扩容后将会生成新的切片
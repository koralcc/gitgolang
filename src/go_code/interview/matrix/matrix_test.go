package matrix

import (
	"fmt"
	"sync"
	"testing"
)

var matrixLength = 2048
var M = 1000000

func BenchmarkMatrixCombination(b *testing.B) {
	//var matrixLength = 1024
	//256*256=65536 //64kb
	//512*512=262144 //256k//本机一级缓存大小,大于一级缓存后，执行的效率明显起来

	matrixA := CreateMatrix(matrixLength)
	matrixB := CreateMatrix(matrixLength)
	//fmt.Println(matrixA)
	b.N = 200
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[i][j] // 	 12	 160191483 ns/op	 6989149 B/op	    1066 allocs/op
			}
		}
	}

}

func BenchmarkMatrixReversedCombination(b *testing.B) {
	//var matrixLength = 1024 //256*256=65536 //64kb
	//512*512=262144 //256k
	b.N = 200
	matrixA := CreateMatrix(matrixLength)
	matrixB := CreateMatrix(matrixLength)
	//fmt.Println(matrixA)
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i] //      3	 352704233 ns/op	27956597 B/op	    4267 allocs/op
			}
		}
	}

}

//嵌套循环最优化，每次循环一小块区域，充分运用一级缓存，本机256kb
func BenchmarkMatrixReversedCombinationPerBlock(b *testing.B) {
	b.N = 200
	//var matrixLength = 1024 //256*256=65536 //64kb
	//512*512=262144 //256k
	blockSize := 64 //缓存行：64Byte
	matrixA := CreateMatrix(matrixLength)
	matrixB := CreateMatrix(matrixLength)
	//fmt.Println(matrixA)
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i += blockSize {
			for j := 0; j < matrixLength; j += blockSize {
				for ii := i; ii < i+blockSize && ii < matrixLength; ii++ {
					for jj := j; jj < j+blockSize && jj < matrixLength; jj++ {
						//fmt.Println(ii, jj)
						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
					}
				}
			}
		}
	}

}

//伪共享
func BenchmarkStructureFalseSharing(b *testing.B) {
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	fmt.Printf("structA:%p\n", &structA)
	fmt.Printf("structB:%p\n", &structB)
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

//内存填充，解决伪共享问题
func BenchmarkStructurePadding(b *testing.B) {
	structA := PaddedStruct{}
	structB := SimpleStruct{}
	fmt.Printf("structA:%p\n", &structA)
	fmt.Printf("structB:%p\n", &structB)
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

type SimpleStruct struct {
	n int
}

type PaddedStruct struct {
	n int
	_ CacheLinePad
}

//内存填充对象
type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

const CacheLinePadSize = 64

//设定缓存行最大值，根据每台电脑cpu不同
//go test -v -bench="." matrix_test.go matrix.go
//-bench="." 运行所有该测试类下的压力测试方法
//go test -v -bench="." -benchmem matrix_test.go matrix.go
//-benchmem 内存分配情况

//windows获取三级缓存信息
//wmic memcache list brief

//https://strikefreedom.top/cpu-caches-theory-and-application

// …or create a new repository on the command line
// echo "# gitgolang" >> README.md
// git init
// git add README.md
// git commit -m "first commit"
// git branch -M main
// git remote add origin https://github.com/koralcc/gitgolang.git
// git push -u origin main

// …or push an existing repository from the command line
// git remote add origin https://github.com/koralcc/gitgolang.git
// git branch -M main
// git push -u origin main

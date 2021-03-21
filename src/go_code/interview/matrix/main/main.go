package main

import (
	"fmt"
	"go_code/interview/matrix"
	"unsafe"
)

func main() {
	var matrixLength = 64000
	matrixA := matrix.CreateMatrix(matrixLength)
	fmt.Println("len", len(matrixA))
	fmt.Println(unsafe.Sizeof(matrixA))
}

package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "abc"))
}

type Stack struct {
	stk []byte
}

func NewStack(bs []byte) *Stack {
	return &Stack{bs}
}

func (this *Stack) Pop() {
	this.stk = this.stk[:(len(this.stk) - 1)]
}

func (this *Stack) Top() byte {
	return this.stk[0]
}

func (this *Stack) Push(b byte) {
	this.stk = append(this.stk, b)

}

func longestCommonSubsequence(text1 string, text2 string) int {
	stack := NewStack([]byte(text1))
	var count = 0
	for j := len(stack.stk) - 1; j > -1; j-- {
		for i := 0; i < len(text2); i++ {
			if stack.Top() != text2[i] {
				continue
			}
			stack.Pop()
			count++
		}
	}
	return count
}

// [1143].最长公共子序列

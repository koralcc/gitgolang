package main

import "fmt"

func main() {

	var list1 = ListNode{10, nil}
	var list2 = ListNode{101, &list1}
	var list3 = ListNode{110, &list2}
	var list4 = ListNode{120, &list3}
	//120 110 101 10
	//101 10 120 110
	ret := rotateRight(&list4, 3)
	fmt.Println(ret.Val)
	fmt.Println(ret.Next.Val)
	fmt.Println(ret.Next.Next.Val)
	fmt.Println(ret.Next.Next.Next.Val)

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	len := len(head)
	if len == 0 {
		return head
	}
	if k >= len {
		k = k % len
	}
	if k == 0 {
		return head
	}
	var endNew = head
	var startNew = head
	//1.将startNew往前移动k个元素
	var i int
	for i < k {
		startNew = startNew.Next
		i++
	}

	//再同时移动endNew和startNew，当startNew到结尾时(startNew.Next == nil)，则得到新的链表头尾
	for startNew.Next != nil {
		startNew = startNew.Next
		endNew = endNew.Next
	}
	//对startNew进行Next指向旧head
	startNew.Next = head
	//对endNew进行next的绑定解除
	ret := endNew.Next
	endNew.Next = nil
	//返回新head
	return ret
}

//首先要计算链表的长度
func len(head *ListNode) int {
	var len int
	for head != nil {
		head = head.Next
		len++
	}
	return len
}

//旋转链表
//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

//f(n) = f(n-1).Next
//分析思路，向右移动k个，因为是链表，所以，从右起第k个即是新头，找到新头

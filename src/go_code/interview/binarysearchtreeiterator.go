package main

import (
	"encoding/json"
	"fmt"
	"go_code/tree"
)

func main() {
	root := tree.GenerateBinaryTree(7, 3, 15, nil, nil, 9, 20)
	str, _ := json.Marshal(root)
	fmt.Println(string(str))

	bst := Constructor(root)
	fmt.Println(bst.ints)
	for bst.HasNext() {
		fmt.Println(bst.Next())
	}
}

type BSTIterator struct {
	ints  []int
	index int
}

func Constructor(root *tree.TreeNode) BSTIterator {
	var bst BSTIterator
	bst.ints = inOrderTraversal(root)
	return bst
}

func (this *BSTIterator) Next() int {
	var currentInt int
	if this.HasNext() {
		currentInt = this.ints[this.index]
		this.index++
	}
	return currentInt
}

func (this *BSTIterator) HasNext() bool {
	if this.index >= len(this.ints) {
		return false
	}
	return true
}

//中序遍历
func inOrderTraversal(root *tree.TreeNode) []int {
	var ret []int
	var inorder func(TN *tree.TreeNode)
	inorder = func(TN *tree.TreeNode) {
		if TN == nil {
			return
		}
		inorder(TN.Left)
		ret = append(ret, TN.Val)
		inorder(TN.Right)
	}
	inorder(root)
	return ret

}

//[173]二叉搜索迭代器

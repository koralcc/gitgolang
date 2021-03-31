package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Inorder(nil)
	tn := generateBinaryTree(1, nil, 3, nil, nil, 4, 5, 6, 7)
	str, _ := json.Marshal(tn)
	fmt.Println(string(str))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Inorder(tn *TreeNode) {
	generateBinaryTree(1, 2, 3, 4, nil)
	if tn == nil {
		return
	}
	Inorder(tn.Left)
	fmt.Println(tn.Val)
	Inorder(tn.Right)

}

//构建二叉树通用方法 1 nil 3 4 5 6 7
func generateBinaryTree(data ...interface{}) (TN *TreeNode) {
	var len = len(data)
	TN = createNode(0, data, len)
	return
}

func createNode(index int, arr []interface{}, len int) (TN *TreeNode) {
	if index > len-1 {
		return
	}
	val, ok := arr[index].(int)
	if !ok {
		return
	}
	TN = &TreeNode{
		Val:   val,
		Left:  createNode(index*2+1, arr, len), //满二叉树左节点是 2*n+1
		Right: createNode(index*2+2, arr, len), //满二叉树右节点是2*n+2
	}
	return
}

//二叉搜索树，中序遍历迭代器实现
//二叉树中序遍历的结果是一串有序序列

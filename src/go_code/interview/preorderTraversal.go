package main

import "fmt"

func main() {
	// tree1 := TreeNode{8, nil, nil}
	// tree2 := TreeNode{7, nil, &tree1}
	// tree3 := TreeNode{11, nil, nil}
	// tree4 := TreeNode{15, &tree2, nil}
	// tree5 := TreeNode{120, &tree3, &tree4}
	//var t TreeNode = nil
	fmt.Println(preorderTraversal(nil))

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//fn = f(n-1).Left
//fn = f(n-1).Right
//1.确定参数和返回值
var retTreeNode []int

func preorderTraversal(root *TreeNode) []int {

	retTreeNode = []int{}
	//2.确定结束条件
	if root == nil {
		return retTreeNode
	}

	//3.确定递归逻辑
	retTreeNode = append(retTreeNode, root.Val)
	//中
	//左
	preorderTraversal(root.Left)
	//右
	preorderTraversal(root.Right)

	return retTreeNode
}

//

//前序遍历
//中-左-右
//左子树遍历完成的条件left、right都等于nil

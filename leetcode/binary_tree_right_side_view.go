package leetcode

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var intSlice []int
	helper(root, &intSlice, 0)
	return intSlice
}

func helper(root *TreeNode, slice *[]int, level int) {
	if root == nil {
		return
	}
	if len(*slice) == level {
		*slice = append(*slice, root.Val)
	}

	helper(root.Right, slice, level+1)
	helper(root.Left, slice, level+1)
}

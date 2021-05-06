package leetcode

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low { // 最低値より小さいとき
		return rangeSumBST(root.Right, low, high) // 右側ノードで再度計算
	}
	if high < root.Val { // 最高値より大きいとき
		return rangeSumBST(root.Left, low, high) // 左側ノードで再度計算
	}

	// rootと左側ノードの該当範囲と右側ノードの該当範囲を足して返す
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

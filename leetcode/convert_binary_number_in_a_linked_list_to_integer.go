package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var decimal int
	for head != nil {
		decimal *= 2 // 位が上がる度に2倍していく
		if head.Val == 1 {
			decimal++
		}
		head = head.Next // 次のノードを渡す
	}

	return decimal
}

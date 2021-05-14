package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	valueMap := map[int]struct{}{}
	firstNode := head // はじめのノード
	for head != nil && head.Next != nil {
		valueMap[head.Val] = struct{}{}
		if _, ok := valueMap[head.Next.Val]; ok { // 同じValがあるとき
			if head.Next.Next != nil {
				head.Next = head.Next.Next
				continue
			} else {
				head.Next = nil
			}
		}
		head = head.Next
	}
	return firstNode
}

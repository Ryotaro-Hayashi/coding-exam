package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// mapで存在確認する
//func deleteDuplicates(head *ListNode) *ListNode {
//	valueMap := map[int]struct{}{}
//	firstNode := head // はじめのノード
//	for head != nil && head.Next != nil {
//		valueMap[head.Val] = struct{}{}
//		if _, ok := valueMap[head.Next.Val]; ok { // 同じValがあるとき
//			if head.Next.Next != nil {
//				head.Next = head.Next.Next
//				continue
//			} else {
//				head.Next = nil
//			}
//		}
//		head = head.Next
//	}
//	return firstNode
//}

func deleteDuplicates(head *ListNode) *ListNode {
	firstNode := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}

	return firstNode
}

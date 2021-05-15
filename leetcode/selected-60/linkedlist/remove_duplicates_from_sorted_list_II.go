package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func deleteDuplicates(head *ListNode) *ListNode {
//	dummyNode := &ListNode{
//		Next: head,
//	}
//	for head != nil && head.Next != nil {
//		if head.Val == head.Next.Val { // 重複しているノードがある
//			for head.Next != nil && head.Val == head.Next.Val { // 重複が終わるノードへ繋ぐ
//				head.Next = head.Next.Next
//			}
//		}
//		head = head.Next
//	}
//
//	return dummyNode.Next
//}

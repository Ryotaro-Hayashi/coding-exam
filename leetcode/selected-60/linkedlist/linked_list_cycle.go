package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// mapでサイクルを検出する方法
func hasCycle(head *ListNode) bool {
	nodeMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}

	return false
}

// サイクルがある場合は隣接ノードで値が同じことを利用する方法
//func hasCycle(head *ListNode) bool {
//	var (
//		fast *ListNode = head
//		slow *ListNode = head
//	)
//	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//		if (fast == slow) {
//			return true
//		}
//	}
//	return false
//}

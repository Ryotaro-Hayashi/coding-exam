package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// mapでサイクルを検出する方法
//func hasCycle(head *ListNode) bool {
//	nodeMap := map[*ListNode]struct{}{}
//	for head != nil {
//		if _, ok := nodeMap[head]; ok {
//			return true
//		}
//		nodeMap[head] = struct{}{}
//		head = head.Next
//	}
//
//	return false
//}

// 正攻法
func hasCycle(head *ListNode) bool {
	var (
		fast *ListNode = head // 2つ先の要素をみるポインタ
		slow *ListNode = head // 1つ先の要素をみるポインタ
	)
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if (fast == slow) { // fastとslowが同じになるのはfastがサイクルを回ってきたときなのでtrueを返す
			return true
		}
	}
	return false
}

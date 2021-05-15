package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var listNode, previousNode *ListNode
//
//	sum := toIntHelper(l1) + toIntHelper(l2)
//
//	for sum > 0 {
//		listNode.Val = sum % 10
//
//		sum /= 10
//		previousNode = listNode
//	}
//}
//
//func toIntHelper(list *ListNode) int {
//	var num int
//	for i :=1; list != nil; {
//		num += list.Val * i
//		i *= 10
//		list = list.Next
//	}
//
//	return num
//}

package swordoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	len := 0
	m := make(map[int]*ListNode)
	m[0] = head
	for head.Next != nil {
		len++
		m[len] = head.Next
		head = head.Next
	}
	return m[len-k+1]
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

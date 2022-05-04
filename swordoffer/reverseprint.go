package swordoffer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	list := make([]int, 0, 10000)
	for head.Next != nil {
		list = append([]int{head.Val},list...)
		head = head.Next
	}
	list = append([]int{head.Val},list...)
	head = head.Next
	return list
}

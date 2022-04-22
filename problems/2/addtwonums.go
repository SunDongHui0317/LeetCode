package addtwonums

import (
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v1 := getValue(l1)
	v2 := getValue(l2)
	s := v1 + v2
	return makeNode(s)
}

func getValue(node *ListNode) int {
	v := 0
	base := 0
	for {
		vb := math.Pow10(base)
		v += node.Val * int(vb)
		base++
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return v
}

func makeNode(v int) *ListNode {
	list := make([]int, 0)
	for {
		rem := v % 10
		list = append(list, rem)
		if v < 10 {
			break
		}
		v = (v - rem) / 10
	}

	if len(list) == 0 {
		return nil
	}
	res := &ListNode{
		Val: list[0],
	}
	temp := res
	for i := 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}
	return res
}

func makeNodeByList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	res := &ListNode{
		Val: list[0],
	}
	temp := res
	for i := 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}
	return res
}

func nodeToList(node *ListNode) []int {
	r := make([]int, 0)
	for {
		if node == nil {
			break
		}
		r = append(r, node.Val)
		node = node.Next
	}
	return r
}

func addTwoList(l1, l2 []int) []int {
	minList := l1
	maxList := l2
	if len(l1) > len(l2) {
		minList = l2
		maxList = l1
	}
	r := maxList
	for i := range minList {
		r[i] = maxList[i] + minList[i]
	}
	for i, v := range r {
		if v >= 10 {
			r[i] -= 10
			if i < len(r)-1 {
				r[i+1]++
			} else {
				r = append(r, 1)
			}
		}
	}
	return r
}

func addTwoNumsList(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := nodeToList(l1)
	list2 := nodeToList(l2)
	list := addTwoList(list1, list2)
	return makeNodeByList(list)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, head := 0, l1
	for {
		// 值相加再加进位
		l1.Val += l2.Val + carry
		// 进位
		carry = l1.Val / 10
		// 求余
		l1.Val %= 10
		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			// 将 l2 的链表阶段接到 l1 上
			l1.Next = l2.Next
			break
		}
		// 执行链表中的下一项
		l1 = l1.Next
		l2 = l2.Next
	}

	// 如果前面的进位没有用到
	for carry != 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1.Next.Val += carry
		carry = l1.Next.Val / 10
		l1.Next.Val %= 10
		l1 = l1.Next
	}

	return head
}

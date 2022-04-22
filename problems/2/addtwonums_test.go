package addtwonums

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNums(t *testing.T) {
	node1 := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	node2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	expected := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	result := AddTwoNumbers(&node1, &node2)
	assert.Equal(t, result.Val, expected.Val)
	assert.Equal(t, result.Next.Val, expected.Next.Val)
	assert.Equal(t, result.Next.Next.Val, expected.Next.Next.Val)
}

func TestAddTwoNumsCarry(t *testing.T) {
	node1 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 7,
			},
		},
	}

	node2 := ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	expected := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	result := AddTwoNumbers(&node1, &node2)
	assert.Equal(t, result.Val, expected.Val)
	assert.Equal(t, result.Next.Val, expected.Next.Val)
	assert.Equal(t, result.Next.Next.Val, expected.Next.Next.Val)
	assert.Equal(t, result.Next.Next.Next.Val, expected.Next.Next.Next.Val)
}

func TestAddTwoNumsMax(t *testing.T) {
	list1 := []int{
		1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
	}
	list2 := []int{
		5,6,4,
	}
	expectedList := []int{
		6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
	}
	node1 := makeNodeByList(list1)
	node2 := makeNodeByList(list2)
	expectedNode := makeNodeByList(expectedList)
	assert.Equal(t, addTwoNumbers(node1, node2), expectedNode)
}

func TestNodeToList(t *testing.T) {
	list1 := []int{
		1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
	}
	node1 := makeNodeByList(list1)
	list := nodeToList(node1)
	assert.Equal(t, list1, list)
}

func TestAddTwoList(t *testing.T) {
	list1 := []int{
		9,8,
	}
	list2 := []int{
		9,8,
	}
	expectedList := []int{
		8,7,1,
	}
	act := addTwoList(list1,list2)
	fmt.Println("act: ", act)
	assert.Equal(t, expectedList, act)
}

func TestAddTwoNumsList(t *testing.T) {
	list1 := []int{
		1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
	}
	list2 := []int{
		5,6,4,
	}
	expectedList := []int{
		6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
	}
	node1 := makeNodeByList(list1)
	node2 := makeNodeByList(list2)
	expectedNode := makeNodeByList(expectedList)
	assert.Equal(t, addTwoNumsList(node1, node2), expectedNode)
}

func TestAddTwoNumsWithoutCarry(t *testing.T) {
	list1 := []int{
		6,5,5,
	}
	list2 := []int{
		5,6,7,
	}
	expectedList := []int{
		1,2,3,
	}
	node1 := makeNodeByList(list1)
	node2 := makeNodeByList(list2)
	expectedNode := makeNodeByList(expectedList)
	assert.Equal(t, addTwoNumbers(node1, node2), expectedNode)
}
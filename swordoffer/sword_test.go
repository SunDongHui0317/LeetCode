package swordoffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	s := []int{
		2, 3, 1, 0, 2, 5, 3,
	}
	assert.Equal(t, 2, findRepeatNumber(s))
}

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{
			1, 4, 7, 11, 15,
		},
		{
			2, 5, 8, 12, 19,
		},
		{
			3, 6, 9, 16, 22,
		},
		{
			10, 13, 14, 17, 24,
		},
		{
			18, 21, 23, 26, 30,
		},
	}
	assert.Equal(t, true, findNumberIn2DArray(matrix, 5))
	assert.Equal(t, false, findNumberIn2DArray(matrix, 20))
}

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	expected :="We%20are%20happy."
	assert.Equal(t, expected, replaceSpace(s))
}

func TestReversePrint(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}
	expected := []int{
		2,3,1,
	}
	assert.Equal(t, expected, reversePrint(&head))
}

func TestBuildTree(t *testing.T) {
	pre := []int{
		3,9,20,15,7,
	}
	in := []int{
		9,3,15,20,7,
	}
	expected := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	assert.Equal(t, &expected, buildTree(pre,in))
}
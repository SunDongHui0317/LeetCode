package swordoffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected := "We%20are%20happy."
	assert.Equal(t, expected, replaceSpace(s))
}

func TestReversePrint(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	expected := []int{
		2, 3, 1,
	}
	assert.Equal(t, expected, reversePrint(&head))
}

func TestBuildTree(t *testing.T) {
	pre := []int{
		3, 9, 20, 15, 7,
	}
	in := []int{
		9, 3, 15, 20, 7,
	}
	expected := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
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
	assert.Equal(t, &expected, buildTree(pre, in))
}

func TestFib(t *testing.T) {
	assert.Equal(t, 3, fib(4))
}

func TestNumWays1(t *testing.T) {
	assert.Equal(t, 2, numWays(2))
}

func TestNumWays2(t *testing.T) {
	assert.Equal(t, 21, numWays(7))
}

func TestNumWays3(t *testing.T) {
	assert.Equal(t, 1, numWays(0))
}

func TestMinArray(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 1, minArray(numbers))
}

func TestMinArray1(t *testing.T) {
	numbers := []int{3, 3, 3, 3, 3}
	assert.Equal(t, 3, minArray(numbers))
}

func TestExist(t *testing.T) {
	board := [][]byte{
		{
			'A', 'B', 'C', 'E',
		}, {
			'S', 'F', 'C', 'S',
		}, {
			'A', 'D', 'E', 'E',
		},
	}
	word := "ABCCED"
	assert.Equal(t, true, exist(board, word))
}

func TestExist1(t *testing.T) {
	board := [][]byte{
		{
			'a', 'b',
		}, {
			'c', 'd',
		},
	}
	word := "abcd"
	assert.Equal(t, false, exist(board, word))
}

func TestMovingCount(t *testing.T) {
	expected := 3
	assert.Equal(t, expected, movingCount(2, 3, 1))
}

func TestMovingCount1(t *testing.T) {
	expected := 1
	assert.Equal(t, expected, movingCount(3, 1, 0))
}

func TestCuttingRope(t *testing.T) {
	expected := 1
	assert.Equal(t, expected, cuttingRope2(2))
}

func TestCuttingRope1(t *testing.T) {
	expected := 36
	assert.Equal(t, expected, cuttingRope2(10))
}

func TestCuttingRope2(t *testing.T) {
	expected := 36
	assert.Equal(t, expected, cuttingRope2(10))
}

func TestHammingWeight(t *testing.T) {
	assert.Equal(t, 31, hammingWeight(4294967293))
}

func TestMyPow(t *testing.T) {
	t.Log(myPow(10.0, 5))
}

func TestPrintNumbers(t *testing.T) {
	expected := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	assert.Equal(t, true, StringSliceEqualBCE(expected, printNumbers(1)))
	//t.Log(printNumbers(1))
}

func StringSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber("0"))
}

func TestIsNumber1(t *testing.T) {
	assert.Equal(t, false, isNumber("e"))
}

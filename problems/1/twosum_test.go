package twosum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{
		2,7,11,15,
	}
	target := 22
	expected := []int{
		1,3,
	}
	assert.Equal(t, twoSum(nums, target), expected)
}

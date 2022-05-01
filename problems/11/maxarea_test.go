package maxarea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea1(t *testing.T) {
	arr := []int{
		1,2,3,
	}
	assert.Equal(t, 2,maxArea(arr))
}

func TestMaxArea2(t *testing.T) {
	arr := []int{
		1,8,6,2,5,4,8,3,7,
	}
	assert.Equal(t, 49,maxArea(arr))
}


func TestMaxArea3(t *testing.T) {
	arr := []int{
		1,1,
	}
	assert.Equal(t, 1,maxArea(arr))
}

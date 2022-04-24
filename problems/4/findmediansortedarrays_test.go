package findmediansortedarrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	l1 := []int {
		1,2,3,
	}
	l2 := []int {
		1,2,
	}
	assert.Equal(t, float64(2), findMedianSortedArrays(l1,l2))
}


func TestFindMedianSortedArrays2(t *testing.T) {
	l1 := []int {
		1,
	}
	l2 := []int {
		1,2,4,
	}
	assert.Equal(t, float64(1.5), findMedianSortedArrays2(l1,l2))
}


func TestFindMedianSortedArrays3(t *testing.T) {
	l1 := []int {
		1,2,
	}
	l2 := []int {
		1,2,5,10,100,
	}
	assert.Equal(t, float64(2), findMedianSortedArrays3(l1,l2))
}

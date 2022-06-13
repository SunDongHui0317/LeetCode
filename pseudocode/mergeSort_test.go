package pseudocode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{
		1,9, 8, 7, 6,
	}
	expected := []int {
		1,6,7,8,9,
	}
	r := mergeSort(a)
	assert.Equal(t, expected, r)
}

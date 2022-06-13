package pseudocode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{
		9,8,7,6,
	}
	expected := []int {
		6,7,8,9,
	}
	r := insertionSort(a)
	assert.Equal(t, expected, r)
}

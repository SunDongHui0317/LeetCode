package reverse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse1(t *testing.T) {
	expected := 321
	assert.Equal(t, expected, reverse(123))
}


func TestReverse2(t *testing.T) {
	expected := -321
	assert.Equal(t, expected, reverse(-123))
}


func TestReverse3(t *testing.T) {
	expected := 21
	assert.Equal(t, expected, reverse(120))
}


func TestReverse4(t *testing.T) {
	expected := 0
	assert.Equal(t, expected, reverse(0))
}

func TestReverse5(t *testing.T) {
	expected := 1
	assert.Equal(t, expected, reverse(10))
}


func TestReverse6(t *testing.T) {
	expected := 0
	assert.Equal(t, expected, reverse100(1534236469))
}



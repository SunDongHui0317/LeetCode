package multiply

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	num1 := "2"
	num2 := "3"
	expected := "6"
	assert.Equal(t, expected, multiply(num1, num2))
}

func TestMultiply1(t *testing.T) {
	num1 := "123"
	num2 := "456"
	expected := "56088"
	assert.Equal(t, expected, multiply(num1, num2))
}

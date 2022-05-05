package ispalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	n := 121
	assert.Equal(t, true, isPalindrome(n))
}


func TestIsPalindrome1(t *testing.T) {
	n := -121
	assert.Equal(t, false, isPalindrome(n))
}


func TestIsPalindrome2(t *testing.T) {
	n := 10
	assert.Equal(t, false, isPalindrome(n))
}

func TestIsPalindrome3(t *testing.T) {
	n := 123123
	assert.Equal(t, false, isPalindrome(n))
}

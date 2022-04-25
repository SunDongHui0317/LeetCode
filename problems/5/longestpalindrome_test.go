package longestpalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome1(t *testing.T) {
	str := "babad"
	expected := "bab"
	assert.Equal(t, expected, longestPalindrome(str))
}


func TestLongestPalindrome2(t *testing.T) {
	str := "cbbd"
	expected := "bb"
	assert.Equal(t, expected, longestPalindrome(str))
}

func TestLongestPalindrome3(t *testing.T) {
	str := "ccc"
	expected := "ccc"
	assert.Equal(t, expected, longestPalindrome(str))
}

func TestLongestPalindrome4(t *testing.T) {
	str := "aacabdkacaa"
	expected := "aca"
	assert.Equal(t, expected, longestPalindrome(str))
}

func TestLongestPalindrome5(t *testing.T) {
	str := "dbaabaabaabd"
	expected := "dbaabaabaabd"
	assert.Equal(t, expected, longestPalindrome(str))
}



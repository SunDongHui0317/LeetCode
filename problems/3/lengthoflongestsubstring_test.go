package lengthoflongestsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	m := "aab"
	assert.Equal(t, lengthOfLongestSubstring(m), 2)
}

func Test2(t *testing.T) {
	m := "dvdf"
	assert.Equal(t, lengthOfLongestSubstring(m), 3)
}

func Test3(t *testing.T) {

	m := "abcabc"
	assert.Equal(t, lengthOfLongestSubstring(m), 3)
}

func Test4(t *testing.T) {
	m := "bbbbb"
	assert.Equal(t, lengthOfLongestSubstring(m), 1)
}

func Test5(t *testing.T) {
	m := "pwwkew"
	assert.Equal(t, lengthOfLongestSubstring(m), 3)
}

func Test6(t *testing.T) {
	m := "c"
	assert.Equal(t, lengthOfLongestSubstring(m), 1)
}

func Test7(t *testing.T) {
	m := "c"
	assert.Equal(t, LengthOfLongestSubstring(m), 1)
}

func Test8(t *testing.T) {
	m := "pwwkew"
	assert.Equal(t, 3, LengthOfLongestSubstring(m))
}

func Test9(t *testing.T) {
	m := "aab"
	assert.Equal(t, 2, LengthOfLongestSubstring(m))
}

func Test10(t *testing.T) {

	m := "abcabc"
	assert.Equal(t,  3, LengthOfLongestSubstring(m))
}
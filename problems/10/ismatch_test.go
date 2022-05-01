package ismatch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	s := "aa"
	p := "a"
	assert.Equal(t, false, isMatch(s, p))
}

func TestIsMatch1(t *testing.T) {
	s := "aa"
	p := "a*"
	assert.Equal(t, true, isMatch(s, p))
}

func TestIsMatch2(t *testing.T) {
	s := "aa"
	p := "a."
	assert.Equal(t, true, isMatch(s, p))
}

func TestIsMatch3(t *testing.T) {
	s := "ab"
	p := ".*"
	assert.Equal(t, true, isMatch(s, p))
}

func TestIsMatch4(t *testing.T) {
	s := "abcdefg"
	p := ".*"
	assert.Equal(t, true, isMatch(s, p))
}

func TestIsMatch5(t *testing.T) {
	s := "abcdefghijklmnopq"
	p := ".bcd*jkk"
	assert.Equal(t, false, isMatch(s, p))
}

func TestIsMatch6(t *testing.T) {
	s := "abcdefghijklmnopq"
	p := ".bcd*jkl*"
	assert.Equal(t, false, isMatch(s, p))
}

package romantoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	s := "III"
	assert.Equal(t, 3, romanToInt(s))
}

func TestRomanToInt1(t *testing.T) {
	s := "IV"
	assert.Equal(t, 4, romanToInt(s))
}

func TestRomanToInt2(t *testing.T) {
	s := "IX"
	assert.Equal(t, 9, romanToInt(s))
}

func TestRomanToInt3(t *testing.T) {
	s := "LVIII"
	assert.Equal(t, 58, romanToInt(s))
}

func TestRomanToInt4(t *testing.T) {
	s := "MCMXCIV"
	assert.Equal(t, 1994, romanToInt(s))
}

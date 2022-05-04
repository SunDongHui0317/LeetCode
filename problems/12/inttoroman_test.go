package inttoroman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "XII", intToRoman(12))
}

func TestIntToRoman1(t *testing.T) {
	assert.Equal(t, "IX", intToRoman(9))
}

func TestIntToRoman2(t *testing.T) {
	assert.Equal(t, "MCCXXIX", intToRoman(1229))
}

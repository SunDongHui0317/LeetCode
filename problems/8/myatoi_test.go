package myatoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := "  42"
	expected := 42
	assert.Equal(t, expected, myAtoi(s))
}


func TestMyAtoi1(t *testing.T) {
	s := "  -42"
	expected := -42
	assert.Equal(t, expected, myAtoi(s))
}

func TestMyAtoi2(t *testing.T) {
	s := "4193 with words "
	expected := 4193
	assert.Equal(t, expected, myAtoi(s))
}

func TestMyAtoi3(t *testing.T) {
	s := "1"
	expected := 1
	assert.Equal(t, expected, myAtoi(s))
}

func TestMyAtoi4(t *testing.T) {
	s := "21474836460"
	expected := 2147483647
	assert.Equal(t, expected, myAtoi(s))
}

func TestMyAtoi5(t *testing.T) {
	s := ""
	expected := 0
	assert.Equal(t, expected, MyAtoi(s))
}


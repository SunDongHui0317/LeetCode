package swordoffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, true, IsMatch("tessssssbssss", "t.s*.b.*"))
}

func TestIsMatch2(t *testing.T) {
	assert.Equal(t, true, IsMatch("teab", "t.a."))
}

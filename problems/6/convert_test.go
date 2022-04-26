package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	str := "PAYPALISHIRING"
	num := 3
	expected := "PAHNAPLSIIGYIR"
	assert.Equal(t, expected, convert(str, num))
}
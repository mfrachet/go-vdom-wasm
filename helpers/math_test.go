package vnh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathMax(t *testing.T) {
	min := 1
	max := 2

	assert.Equal(t, max, Max(min, max))
	assert.Equal(t, max, Max(max, min))
}

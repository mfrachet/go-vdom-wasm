package vnh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotNil(t *testing.T) {

	str := "Hello world"
	pointer := &str

	assert.Equal(t, false, NotNil(nil))
	assert.Equal(t, true, NotNil(pointer))
}

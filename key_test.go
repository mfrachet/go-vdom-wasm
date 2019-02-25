package vn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

import vn "github.com/mfrachet/go-vdom-wasm"

func TestKey_Accessors(t *testing.T) {
	key := vn.Key{}
	key.SetValue("Hello")

	assert.Equal(t, "Hello", key.GetValue())
}

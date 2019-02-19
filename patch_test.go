package vn

import (
	"fmt"
	"testing"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	"github.com/stretchr/testify/assert"
)

func TestPatchInitialize(t *testing.T) {
	doc := vnd.GetDocument()

	fmt.Println("lol", doc.GetBinding())

	assert.Equal(t, true, false)
}

package vn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttrs_Sanitize(t *testing.T) {
	attrs := &Attrs{}

	sanitized := Sanitize(attrs)

	assert.Equal(t, false, sanitized.Props == nil)
	assert.Equal(t, false, sanitized.Events == nil)
}

func TestAttrs_Sanitize_No_Changes(t *testing.T) {
	attrs := &Attrs{Props: &Props{}, Events: &Ev{}}

	sanitized := Sanitize(attrs)

	assert.Equal(t, *attrs, sanitized)
}

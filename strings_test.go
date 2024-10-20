package golodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	is := assert.New(t)

	ret := Substring("hello world", 3, 3)
	expect := "lo "

	is.Equal(expect, ret)
}

func TestCapitalize(t *testing.T) {
	is := assert.New(t)

	ret := Capitalize("hello world")
	expect := "Hello World"

	is.Equal(expect, ret)
}

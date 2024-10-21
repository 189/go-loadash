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

// PascalCase converts string to pascal case.
func TestPascalCase(t *testing.T) {
	is := assert.New(t)

	ret := PascalCase("hello world")
	expect := "HelloWorld"

	is.Equal(expect, ret)

}

// CamelCase converts string to camel case.
func TestCamelCase(t *testing.T) {
	is := assert.New(t)

	ret := CamelCase("hello world")
	expect := "helloWorld"

	is.Equal(expect, ret)
}

// KebabCase converts string to kebab case.
func TestKebabCase(t *testing.T) {
	is := assert.New(t)

	ret := KebabCase("hello world")
	expect := "hello-world"

	is.Equal(expect, ret)
}

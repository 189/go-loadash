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

func TestPascalCase(t *testing.T) {
	is := assert.New(t)

	ret := PascalCase("hello world")
	expect := "HelloWorld"

	is.Equal(expect, ret)

}

func TestCamelCase(t *testing.T) {
	is := assert.New(t)

	ret := CamelCase("hello world")
	expect := "helloWorld"

	is.Equal(expect, ret)
}

func TestKebabCase(t *testing.T) {
	is := assert.New(t)

	ret := KebabCase("hello world")
	expect := "hello-world"

	is.Equal(expect, ret)
}

func TestIsAlpha(t *testing.T) {
	is := assert.New(t)

	rst := IsAlpha("13131")
	is.Equal(false, rst)

	rst = IsAlpha("abcdef")
	is.Equal(true, rst)
}

package golodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBignumberMult(t *testing.T) {
	is := assert.New(t)

	v, _ := BignumberMult("0.001", "1000000000000000000")
	is.Equal("1000000000000000", v)

	v, _ = BignumberMult("0.001", "1000000000000000000", "0.1")
	is.Equal("100000000000000", v)

	v, _ = BignumberMult("0.001", "1000000000000000000", "0.1", "10")
	is.Equal("1000000000000000", v)
}

func TestBignumberPlus(t *testing.T) {
	is := assert.New(t)

	v, _ := BignumberPlus("0.1", "0.2")
	is.Equal("0.3", v)

	v, _ = BignumberPlus("0.1", "0.2", "0.3")
	is.Equal("0.6", v)

	_, err := BignumberPlus("0.1")
	is.Error(err)
}

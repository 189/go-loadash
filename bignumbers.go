package golodash

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// BignumberMult return the product of mutiple big numbers
func BignumberMult(input string, rest ...string) (string, error) {
	if len(rest) == 0 {
		return "", fmt.Errorf("more multiplicands are needed")
	}

	total, err := decimal.NewFromString(input)
	if err != nil {
		return "", err
	}

	for _, v := range rest {
		v2, err := decimal.NewFromString(v)
		if err != nil {
			return "", err
		}

		total = total.Mul(v2)
	}

	return total.String(), nil
}

// BignumberPlus return the sum of multiple big/decimal numbers, BignumberPlus("0.1", "0.2") = 0.3,
func BignumberPlus(input string, rest ...string) (string, error) {
	if len(rest) == 0 {
		return "", fmt.Errorf("more multiplicands are needed")
	}

	total, err := decimal.NewFromString(input)
	if err != nil {
		return "", err
	}

	for _, v := range rest {
		v2, err := decimal.NewFromString(v)
		if err != nil {
			return "", err
		}

		total = total.Add(v2)
	}

	return total.String(), nil
}

package golodash

import "fmt"

func ExmapleBignumberMult() {
	v, _ := BignumberMult("0.001", "10", "1")
	fmt.Println(v)
	// Output:
	// "0.01"
}

func ExmapleBignumberPlus() {
	v, _ := BignumberPlus("0.1", "0.2", "0.3")
	fmt.Println(v)
	// Output:
	// "0.6"
}

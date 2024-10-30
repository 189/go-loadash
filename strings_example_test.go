package golodash

import "fmt"

func ExampleIsAlpha() {
	ret := IsAlpha("basdcf")
	fmt.Println(ret)
	// Output
	// true

}

func ExampleIsAlphaNumeric() {
	rst := IsAlphaNumeric("abcd13")
	fmt.Println(rst)
	// Output:
	// true
}

func ExampleIsNumeric() {
	rst := IsNumeric("1234")
	fmt.Println(rst)
	// Output:
	// true
}

func ExamplePadLeft() {
	rst := PadLeft("x", 5)
	fmt.Println(rst)
	// Output:
	// 0000x
}

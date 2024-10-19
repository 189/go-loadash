package golodash

import (
	"fmt"
)

func ExampleReduce() {
	nums := []int32{1, 2, 3, 4, 5}
	total := Reduce(nums, 0, func(prev int32, now int32) int32 {
		return prev + now
	})
	fmt.Printf("%v", total)
	// Output:
	// 15
}

func ExampleMap() {
	nums := []int32{1, 2, 3, 4, 5}
	rst := Map(nums, func(item int32) int32 {
		return item * 2
	})

	fmt.Printf("%v", rst)
	// Output:
	// [2 4 6 8 10]
}

func ExampleFilter() {
	nums := []int32{1, 2, 3, 4, 5}
	rst := Filter(nums, func(item int32) bool {
		return item > 4
	})

	fmt.Printf("%v", rst)
	// Output:
	// [5]
}

func ExampleFlat() {
	items := [][]int{
		{1, 2, 3},
		{1},
		{4, 5, 6},
	}
	rst := Flat(items)
	fmt.Printf("%v", rst)
	// Output:
	// [1 2 3 1 4 5 6]
}

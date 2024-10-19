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
	// 16
}

func ExampleMap() {
	nums := []int32{1, 2, 3, 4, 5}
	ret := Map(nums, func(item int32) int32 {
		return item * 2
	})

	fmt.Printf("%v", ret)
	// Output:
	// {2,4,6,8,10}
}

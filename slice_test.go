package golodash

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	nums := []int32{1, 2, 3, 4, 5}
	total := Reduce(nums, 0, func(prev int32, now int32) int32 {
		return prev + now
	})
	expected := 15
	if total != int32(expected) {
		t.Fatalf(`expect %d, but got %d`, expected, total)
	}
}

func TestMap(t *testing.T) {
	nums := []int32{1, 2, 3, 4, 5}
	ret := Map(nums, func(item int32) int32 {
		return item * 2
	})

	expected := []int32{2, 4, 6, 8, 10}
	if !slices.Equal(expected, ret) {
		t.Fatalf(`expect %v, but got %v`, expected, ret)
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	nums := []int32{1, 2, 3, 4, 5}
	ret := Filter(nums, func(item int32) bool {
		return item > 4
	})
	expected := []int32{5}

	is.Equal(expected, ret)
}

func TestSome(t *testing.T) {
	nums := []int32{1, 2, 3, 4}

	if rst := Some(nums, func(item int32) bool { return item == 2 }); !rst {
		t.Errorf("expect true, but got false")
		return
	}

	if rst := Some(nums, func(item int32) bool { return item == 5 }); rst {
		t.Errorf("expect false, but got true")
		return
	}
}

func TestEvery(t *testing.T) {
	nums := []int32{1, 2, 3, 4}

	if rst := Every(nums, func(item int32) bool { return item > 0 }); !rst {
		t.Errorf("expect true, but got false")
		return
	}

	if rst := Every(nums, func(item int32) bool { return item > 2 }); rst {
		t.Errorf("expect false, but got true")
		return
	}
}

func TestForEach(t *testing.T) {
	nums := []int32{1, 2, 3, 4}
	count := 0
	total := len(nums)

	ForEach(nums, func(idx int, item int32) bool {
		count++
		return true
	})
	if total != count {
		t.Errorf("expect to loop %d times, but got %d", total, count)
		return
	}

	count = 0
	expected := 2

	ForEach(nums, func(idx int, item int32) bool {
		count++
		return item != 2
	})
	if count != expected {
		t.Errorf("expect to loop %d times, but got %d", expected, count)
		return
	}
}

func TestFlat(t *testing.T) {
	t.Parallel()

	is := assert.New(t)
	items := [][]int{
		{1, 2, 3},
		{1},
		{4, 5, 6},
	}
	rst := Flat(items)
	expected := []int{
		1, 2, 3, 1, 4, 5, 6,
	}
	is.Equal(rst, expected)
}

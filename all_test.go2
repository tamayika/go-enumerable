package enumerable

import "testing"

func AllTest(t *testing.T) {
	slice := Slice([]int{10, 20, 30})
	equal(t, true, All(slice, func(v int) bool {
		return v > 0
	}))
	equal(t, false, All(slice, func(v int) bool {
		return v > 10
	}))
	empty := Slice([]int{})
	equal(t, true, All(empty, func(v int) bool {
		return v > 10
	}))
}

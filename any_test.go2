package enumerable

import "testing"

func AnyTest(t *testing.T) {
	slice := Slice([]int{10, 20, 30})
	equal(t, true, Any(slice, func(v int) bool {
		return v > 10
	}))
	equal(t, false, Any(slice, func(v int) bool {
		return v > 30
	}))
	empty := Slice([]int{})
	equal(t, false, Any(empty, func(v int) bool {
		return v > 10
	}))
}

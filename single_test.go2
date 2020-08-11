package enumerable

import "testing"

func TestSingle(t *testing.T) {
	equal(t, 10, Single(Slice([]int{10})))
	expectPanic(t, func() { Single(Slice([]int{})) })
	expectPanic(t, func() { Single(Slice([]int{10, 10})) })
}

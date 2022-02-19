package enumerable

import "testing"

func TestFirst(t *testing.T) {
	equal(t, 10, First(Slice([]int{10})))
	expectPanic(t, func() { First(Slice([]int{})) })
}

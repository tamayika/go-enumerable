package enumerable

import "testing"

func TestMin(t *testing.T) {
	equal(t, 10, Min(Slice([]int{10, 20, 30})))
	expectPanic(t, func() { Min(Slice([]int{})) })
}

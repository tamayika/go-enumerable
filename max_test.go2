package enumerable

import "testing"

func TestMax(t *testing.T) {
	equal(t, 30, Max(Slice([]int{10, 20, 30})))
	expectPanic(t, func() { Max(Slice([]int{})) })
}

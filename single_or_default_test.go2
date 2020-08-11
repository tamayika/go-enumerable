package enumerable

import "testing"

func TestSingleOrDefault(t *testing.T) {
	equal(t, 10, SingleOrDefault(Slice([]int{10})))
	equal(t, 0, SingleOrDefault(Slice([]int{})))
	expectPanic(t, func() { SingleOrDefault(Slice([]int{10, 10})) })
}

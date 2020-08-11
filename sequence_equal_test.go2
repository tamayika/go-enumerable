package enumerable

import "testing"

func TestSequenceEqual(t *testing.T) {
	equal(t, true, SequenceEqual(Slice([]int{}), Slice([]int{})))
	equal(t, false, SequenceEqual(Slice([]int{10}), Slice([]int{})))
	equal(t, false, SequenceEqual(Slice([]int{}), Slice([]int{10})))
	equal(t, true, SequenceEqual(Slice([]int{10}), Slice([]int{10})))
	equal(t, false, SequenceEqual(Slice([]int{10, 10}), Slice([]int{10, 20})))
	equal(t, true, SequenceEqual(Slice([]int{10, 20}), Slice([]int{10, 20})))
}

package enumerable

import "testing"

func TestToSlice(t *testing.T) {
	equal(t, []int{}, ToSlice(Slice([]int{})))
	equal(t, []int{10}, ToSlice(Slice([]int{10})))
	equal(t, []int{10, 20}, ToSlice(Slice([]int{10, 20})))
}

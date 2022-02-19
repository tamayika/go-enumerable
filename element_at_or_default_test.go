package enumerable

import "testing"

func TestElementAtOrDefault(t *testing.T) {
	equal(t, 10, ElementAtOrDefault(Slice([]int{10, 20}), 0))
	equal(t, 20, ElementAtOrDefault(Slice([]int{10, 20}), 1))
	equal(t, 0, ElementAtOrDefault(Slice([]int{10, 20}), 2))
}

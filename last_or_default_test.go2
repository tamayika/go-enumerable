package enumerable

import "testing"

func TestLastOrDefault(t *testing.T) {
	equal(t, 20, LastOrDefault(Slice([]int{10, 20})))
	equal(t, 0, LastOrDefault(Slice([]int{})))
}

package enumerable

import "testing"

func TestFirstOrDefault(t *testing.T) {
	equal(t, 10, FirstOrDefault(Slice([]int{10})))
	equal(t, 0, FirstOrDefault(Slice([]int{})))
}

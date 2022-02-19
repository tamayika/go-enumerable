package enumerable

import "testing"

func TestSum(t *testing.T) {
	equal(t, 0, Sum(Slice([]int{})))
	equal(t, 30, Sum(Slice([]int{10, 20})))
}

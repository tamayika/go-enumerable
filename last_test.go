package enumerable

import "testing"

func TestLast(t *testing.T) {
	equal(t, 20, Last(Slice([]int{10, 20})))
	expectPanic(t, func() { Last(Slice([]int{})) })
}

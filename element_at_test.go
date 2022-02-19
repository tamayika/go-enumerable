package enumerable

import "testing"

func TestElementAt(t *testing.T) {
	equal(t, 10, ElementAt(Slice([]int{10, 20}), 0))
	equal(t, 20, ElementAt(Slice([]int{10, 20}), 1))
	expectPanic(t, func() { ElementAt(Slice([]int{10, 20}), 2) })
}

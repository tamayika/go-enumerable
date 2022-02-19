package enumerable

import "testing"

func TestWhere(t *testing.T) {
	enumerator := Where(Slice([]int{10, 20, 30}), func(v int) bool {
		return v > 10
	}).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestTakeWhile(t *testing.T) {
	enumerator := TakeWhile(Slice([]int{10, 20, 30}), func(v int) bool {
		return v < 30
	}).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

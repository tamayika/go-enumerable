package enumerable

import "testing"

func TestReverse(t *testing.T) {
	enumerator := Reverse(Slice([]int{1, 2, 3, 4})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 4, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 3, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 2, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 1, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

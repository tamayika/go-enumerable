package enumerable

import "testing"

func TestTake(t *testing.T) {
	enumerator := Take(Slice([]int{10, 20}), 0).Enumerator()
	equal(t, false, enumerator.MoveNext())
	enumerator = Take(Slice([]int{10, 20}), 1).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
	enumerator = Take(Slice([]int{10, 20}), 3).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestDefaultIfEmpty(t *testing.T) {
	enumerator := DefaultIfEmpty(Slice([]int{})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 0, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
	enumerator = DefaultIfEmpty(Slice([]int{10, 20})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

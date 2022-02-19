package enumerable

import "testing"

func TestUnion(t *testing.T) {
	enumerator := Union(
		Slice([]int{5, 3, 9, 7, 5, 9, 3, 7}),
		Slice([]int{8, 3, 6, 4, 4, 9, 1, 0}),
	).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 5, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 3, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 9, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 7, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 8, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 6, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 4, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 1, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 0, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestOrderByDescending(t *testing.T) {
	enumerator := OrderByDescending(Slice([]int{20, 30, 10})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

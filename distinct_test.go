package enumerable

import "testing"

func TestDistinct(t *testing.T) {
	enumerator := Distinct(Slice([]int{10, 10, 20, 20, 30})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

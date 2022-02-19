package enumerable

import "testing"

func TestConcat(t *testing.T) {
	enumerator := Concat(Slice([]int{10, 20}), Slice([]int{30, 40})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 40, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

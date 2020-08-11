package enumerable

import "testing"

func TestExcept(t *testing.T) {
	enumerator := Except(Slice([]int{10, 20, 30, 40}), Slice([]int{0, 20})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 40, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestPrepend(t *testing.T) {
	enumerator := Prepend(Slice([]int{10, 20}), 5).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 5, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestRepat(t *testing.T) {
	enumerator := Repeat(10, 3).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

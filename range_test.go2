package enumerable

import "testing"

func TestRange(t *testing.T) {
	enumerator := Range(5, 3).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 5, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 6, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 7, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

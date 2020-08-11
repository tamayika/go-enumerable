package enumerable

import "testing"

func TestSelect(t *testing.T) {
	enumerator := Select(Slice([]string{"a", "b"}), func(s string) string {
		return s + s
	}).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, "aa", enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, "bb", enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

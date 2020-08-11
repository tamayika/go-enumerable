package enumerable

import "testing"

func TestSkip(t *testing.T) {
	enumerator := Skip(Slice([]int{}), 10).Enumerator()
	equal(t, false, enumerator.MoveNext())
	enumerator = Skip(Slice([]int{10, 20}), 0).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
	enumerator = Skip(Slice([]int{10, 20, 30}), 1).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

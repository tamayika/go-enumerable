package enumerable

import "testing"

func TestSkipLast(t *testing.T) {
	enumerator := SkipLast(Slice([]int{10, 20}), 0).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
	enumerator = SkipLast(Slice([]int{10, 20, 30}), 1).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
	enumerator = SkipLast(Slice([]int{10, 20, 30}), 10).Enumerator()
	equal(t, false, enumerator.MoveNext())
}

package enumerable

import "testing"

func TestSkipWhile(t *testing.T) {
	enumerator := SkipWhile(Slice([]int{}), func(t int) bool {
		return false
	}).Enumerator()
	equal(t, false, enumerator.MoveNext())
	enumerator = SkipWhile(Slice([]int{10, 20, 30}), func(t int) bool {
		return t < 20
	}).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

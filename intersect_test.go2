package enumerable

import "testing"

func TestIntersect(t *testing.T) {
	enumerator := Intersect(Slice([]int{10, 20, 30, 40}), Slice([]int{0, 20})).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

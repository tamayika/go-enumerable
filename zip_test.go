package enumerable

import "testing"

func TestZip(t *testing.T) {
	enumerator := Zip(Slice([]int{10, 20}), Slice([]int{30, 40}), func(x, y int) int { return x + y }).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 40, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 60, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

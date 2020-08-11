package enumerable

import "testing"

func TestGroupBy(t *testing.T) {
	enumerator := GroupBy(Slice([]int{10, 20, 20, 30, 30, 30}),
		func(t int) int { return t },
		func(t int) int { return t },
		func(key int, element int, result int) int {
			return result + 1
		},
	).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 1, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 2, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 3, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

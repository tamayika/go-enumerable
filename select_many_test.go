package enumerable

import "testing"

func testSelectManyCollectionFunc(v int) IEnumerable[int] {
	return Slice([]int{v, v})
}

func testSelectManyResultFunc(v int, collectionElement int) int {
	return collectionElement
}

func TestSelectMany(t *testing.T) {
	enumerator := SelectMany(
		Slice([]int{10, 20, 30}),
		testSelectManyCollectionFunc,
		testSelectManyResultFunc,
	).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 30, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

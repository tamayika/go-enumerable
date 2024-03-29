package enumerable

func SequenceEqual[TSource comparable](first, second IEnumerable[TSource]) bool {
	firstEnumerator := first.Enumerator()
	secondEnumerator := second.Enumerator()
	for {
		firstOK := firstEnumerator.MoveNext()
		secondOK := secondEnumerator.MoveNext()
		if firstOK != secondOK {
			return false
		}
		if !firstOK {
			return true
		}
		if firstEnumerator.Current() != secondEnumerator.Current() {
			return false
		}
	}
}

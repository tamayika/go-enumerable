package enumerable

func Count[TSource any](source IEnumerable[TSource]) int {
	enumerator := source.Enumerator()
	count := 0
	for enumerator.MoveNext() {
		count++
	}
	return count
}

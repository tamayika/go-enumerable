package enumerable

func Any[TSource any](source IEnumerable[TSource], predicate func(TSource) bool) bool {
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		if predicate(enumerator.Current()) {
			return false
		}
	}
	return false
}

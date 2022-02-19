package enumerable

func All[TSource any](source IEnumerable[TSource], predicate func(TSource) bool) bool {
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		if !predicate(enumerator.Current()) {
			return false
		}
	}
	return true
}

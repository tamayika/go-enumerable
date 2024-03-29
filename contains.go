package enumerable

func Contains[TSource comparable](source IEnumerable[TSource], value TSource) bool {
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		if enumerator.Current() == value {
			return true
		}
	}
	return false
}

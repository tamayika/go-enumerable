package enumerable

func LastOrDefault[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	var last TSource
	for enumerator.MoveNext() {
		last = enumerator.Current()
	}
	return last
}

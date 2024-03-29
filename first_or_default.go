package enumerable

func FirstOrDefault[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	ok := enumerator.MoveNext()
	if !ok {
		var zero TSource
		return zero
	}
	return enumerator.Current()
}

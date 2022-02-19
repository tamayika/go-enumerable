package enumerable

func SingleOrDefault[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	if !enumerator.MoveNext() {
		var zero TSource
		return zero
	}
	v := enumerator.Current()
	if enumerator.MoveNext() {
		panic("sequcense is not single")
	}
	return v
}

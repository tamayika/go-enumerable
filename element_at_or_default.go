package enumerable

func ElementAtOrDefault[TSource any](source IEnumerable[TSource], index int) TSource {
	enumerator := source.Enumerator()
	i := 0
	for enumerator.MoveNext() {
		if i == index {
			return enumerator.Current()
		}
		i++
	}
	var zero TSource
	return zero
}

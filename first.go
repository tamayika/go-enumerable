package enumerable

func First[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	ok := enumerator.MoveNext()
	if !ok {
		panic("empty")
	}
	return enumerator.Current()
}

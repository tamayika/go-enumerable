package enumerable

func Single[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	if !enumerator.MoveNext() {
		panic("empty sequence")
	}
	v := enumerator.Current()
	if enumerator.MoveNext() {
		panic("sequcense is not single")
	}
	return v
}

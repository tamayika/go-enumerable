package enumerable

func ElementAt[TSource any](source IEnumerable[TSource], index int) TSource {
	enumerator := source.Enumerator()
	i := 0
	for enumerator.MoveNext() {
		if i == index {
			return enumerator.Current()
		}
		i++
	}
	panic("index is out of enumerable")
}

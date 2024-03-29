package enumerable

func Sum[TSource number](source IEnumerable[TSource]) TSource {
	var sum TSource
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		sum += enumerator.Current()
	}
	return sum
}

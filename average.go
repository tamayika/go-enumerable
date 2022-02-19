package enumerable

func Average[TSource number](source IEnumerable[TSource]) TSource {
	var zero TSource
	var sum TSource
	var count TSource
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		sum += enumerator.Current()
		count++
	}
	if count == zero {
		return zero
	}
	return sum / count
}

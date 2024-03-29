package enumerable

func Last[TSource any](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	var last TSource
	var lastSet bool
	for enumerator.MoveNext() {
		last = enumerator.Current()
		lastSet = true
	}
	if !lastSet {
		panic("empty enumerable")
	}
	return last
}

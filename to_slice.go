package enumerable

func ToSlice[TSource any](source IEnumerable[TSource]) []TSource {
	slice := make([]TSource, 0)
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		slice = append(slice, enumerator.Current())
	}
	return slice
}

package enumerable

func ToMap[TSource comparable, TKey comparable, TElement any](
	source IEnumerable[TSource], keySelector func(TSource) TKey,
	elementSelector func(TSource) TElement) map[TKey]TElement {
	enumerator := source.Enumerator()
	ret := map[TKey]TElement{}
	for enumerator.MoveNext() {
		key := keySelector(enumerator.Current())
		element := elementSelector(enumerator.Current())
		ret[key] = element
	}
	return ret
}

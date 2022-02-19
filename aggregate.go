package enumerable

func Aggregate[TSource, TAccumulate, TResult any](
	source IEnumerable[TSource], seed TAccumulate,
	f func(TAccumulate, TSource) TAccumulate, resultSelector func(TAccumulate) TResult) TResult {
	enumerator := source.Enumerator()
	for enumerator.MoveNext() {
		seed = f(seed, enumerator.Current())
	}
	return resultSelector(seed)
}

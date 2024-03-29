package enumerable

func TakeLast[TSource any](source IEnumerable[TSource], count int) IEnumerable[TSource] {
	return takeLastEnumerable[TSource]{source, count}
}

type takeLastEnumerable[TSource any] struct {
	source IEnumerable[TSource]
	count  int
}

func (t takeLastEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &takeLastEnumerator[TSource]{enumerator: t.source.Enumerator(), count: t.count}
}

type takeLastEnumerator[TSource any] struct {
	enumerator     IEnumerator[TSource]
	count          int
	tookEnumerator IEnumerator[TSource]
}

func (t *takeLastEnumerator[TSource]) Current() TSource {
	return t.tookEnumerator.Current()
}

func (t *takeLastEnumerator[TSource]) MoveNext() bool {
	if t.tookEnumerator == nil {
		var slice []TSource
		for t.enumerator.MoveNext() {
			slice = append(slice, t.enumerator.Current())
		}
		start := len(slice) - t.count
		if start < 0 {
			start = 0
		}
		t.tookEnumerator = &sliceEnumerator[TSource]{slice: slice[start:], cursor: -1}
	}
	return t.tookEnumerator.MoveNext()
}

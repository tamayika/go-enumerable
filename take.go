package enumerable

func Take[TSource any](source IEnumerable[TSource], count int) IEnumerable[TSource] {
	return takeEnumerable[TSource]{source, count}
}

type takeEnumerable[TSource any] struct {
	source IEnumerable[TSource]
	count  int
}

func (t takeEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &takeEnumerator[TSource]{enumerator: t.source.Enumerator(), count: t.count}
}

type takeEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	count      int
	tookCount  int
}

func (t *takeEnumerator[TSource]) Current() TSource {
	return t.enumerator.Current()
}

func (t *takeEnumerator[TSource]) MoveNext() bool {
	if t.tookCount >= t.count {
		return false
	}
	t.tookCount++
	return t.enumerator.MoveNext()
}

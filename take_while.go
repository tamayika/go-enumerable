package enumerable

func TakeWhile[TSource any](source IEnumerable[TSource], predicate func(TSource) bool) IEnumerable[TSource] {
	return takeWhileEnumerable[TSource]{source, predicate}
}

type takeWhileEnumerable[TSource any] struct {
	source    IEnumerable[TSource]
	predicate func(TSource) bool
}

func (t takeWhileEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &takeWhileEnumerator[TSource]{enumerator: t.source.Enumerator(), predicate: t.predicate}
}

type takeWhileEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	predicate  func(TSource) bool
}

func (t *takeWhileEnumerator[TSource]) Current() TSource {
	return t.enumerator.Current()
}

func (t *takeWhileEnumerator[TSource]) MoveNext() bool {
	ok := t.enumerator.MoveNext()
	if !ok {
		return false
	}
	return t.predicate(t.enumerator.Current())
}

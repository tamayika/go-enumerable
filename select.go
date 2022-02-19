package enumerable

func Select[TSource, TResult any](source IEnumerable[TSource], selector func(TSource) TResult) IEnumerable[TResult] {
	return selectEnumerable[TSource, TResult]{source, selector}
}

type selectEnumerable[TSource, TResult any] struct {
	source   IEnumerable[TSource]
	selector func(TSource) TResult
}

func (s selectEnumerable[TSource, TResult]) Enumerator() IEnumerator[TResult] {
	return &selectEnumerator[TSource, TResult]{enumerator: s.source.Enumerator(), selector: s.selector}
}

type selectEnumerator[TSource, TResult any] struct {
	enumerator IEnumerator[TSource]
	selector   func(TSource) TResult
}

func (s *selectEnumerator[TSource, TResult]) Current() TResult {
	return s.selector(s.enumerator.Current())
}

func (s *selectEnumerator[TSource, TResult]) MoveNext() bool {
	return s.enumerator.MoveNext()
}

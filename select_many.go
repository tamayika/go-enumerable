package enumerable

func SelectMany[TSource, TCollection, TResult any](
	source IEnumerable[TSource],
	collectionSelector func(TSource) IEnumerable[TCollection],
	resultSelector func(TSource, TCollection) TResult,
) IEnumerable[TResult] {
	return selectManyEnumerable[TSource, TCollection, TResult]{
		source,
		collectionSelector,
		resultSelector,
	}
}

type selectManyEnumerable[TSource, TCollection, TResult any] struct {
	source             IEnumerable[TSource]
	collectionSelector func(TSource) IEnumerable[TCollection]
	resultSelector     func(TSource, TCollection) TResult
}

func (s selectManyEnumerable[TSource, TCollection, TResult]) Enumerator() IEnumerator[TResult] {
	return &selectManyEnumerator[TSource, TCollection, TResult]{
		enumerator:         s.source.Enumerator(),
		collectionSelector: s.collectionSelector,
		resultSelector:     s.resultSelector,
	}
}

type selectManyEnumerator[TSource, TCollection, TResult any] struct {
	enumerator          IEnumerator[TSource]
	collectionSelector  func(TSource) IEnumerable[TCollection]
	resultSelector      func(TSource, TCollection) TResult
	collectionEumerator IEnumerator[TCollection]
	result              TResult
}

func (s *selectManyEnumerator[TSource, TCollection, TResult]) Current() TResult {
	return s.result
}

func (s *selectManyEnumerator[TSource, TCollection, TResult]) MoveNext() bool {
	for {
		if s.collectionEumerator == nil {
			ok := s.enumerator.MoveNext()
			if !ok {
				return false
			}
			s.collectionEumerator = s.collectionSelector(s.enumerator.Current()).Enumerator()
		}
		ok := s.collectionEumerator.MoveNext()
		if !ok {
			s.collectionEumerator = nil
			continue
		}
		s.result = s.resultSelector(s.enumerator.Current(), s.collectionEumerator.Current())
		return true
	}
}

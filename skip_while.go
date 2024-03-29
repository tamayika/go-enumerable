package enumerable

func SkipWhile[TSource any](source IEnumerable[TSource], predicate func(TSource) bool) IEnumerable[TSource] {
	return skipWhileEnumerable[TSource]{source, predicate}
}

type skipWhileEnumerable[TSource any] struct {
	source    IEnumerable[TSource]
	predicate func(TSource) bool
}

func (s skipWhileEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &skipWhileEnumerator[TSource]{enumerator: s.source.Enumerator(), predicate: s.predicate}
}

type skipWhileEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	predicate  func(TSource) bool
	skipped    bool
}

func (s *skipWhileEnumerator[TSource]) Current() TSource {
	return s.enumerator.Current()
}

func (s *skipWhileEnumerator[TSource]) MoveNext() bool {
	if !s.skipped {
		for {
			ok := s.enumerator.MoveNext()
			if !ok {
				s.skipped = true
				return false
			}
			if !s.predicate(s.enumerator.Current()) {
				s.skipped = true
				return true
			}
		}
	}
	return s.enumerator.MoveNext()
}

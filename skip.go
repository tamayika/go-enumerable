package enumerable

func Skip[TSource any](source IEnumerable[TSource], count int) IEnumerable[TSource] {
	return skipEnumerable[TSource]{source, count}
}

type skipEnumerable[TSource any] struct {
	source IEnumerable[TSource]
	count  int
}

func (s skipEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &skipEnumerator[TSource]{enumerator: s.source.Enumerator(), count: s.count}
}

type skipEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	count      int
	skipped    bool
}

func (s *skipEnumerator[TSource]) Current() TSource {
	return s.enumerator.Current()
}

func (s *skipEnumerator[TSource]) MoveNext() bool {
	if !s.skipped {
		var count int
		for {
			if count == s.count {
				s.skipped = true
				break
			}
			ok := s.enumerator.MoveNext()
			if !ok {
				s.skipped = true
				return false
			}
			count++
		}
	}
	return s.enumerator.MoveNext()
}

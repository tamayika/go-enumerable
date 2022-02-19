package enumerable

func SkipLast[TSource any](source IEnumerable[TSource], count int) IEnumerable[TSource] {
	return skipLastEnumerable[TSource]{source, count}
}

type skipLastEnumerable[TSource any] struct {
	source IEnumerable[TSource]
	count  int
}

func (s skipLastEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &skipLastEnumerator[TSource]{enumerator: s.source.Enumerator(), count: s.count}
}

type skipLastEnumerator[TSource any] struct {
	enumerator        IEnumerator[TSource]
	count             int
	skippedEnumerator IEnumerator[TSource]
}

func (s *skipLastEnumerator[TSource]) Current() TSource {
	return s.skippedEnumerator.Current()
}

func (s *skipLastEnumerator[TSource]) MoveNext() bool {
	if s.skippedEnumerator == nil {
		var slice []TSource
		for s.enumerator.MoveNext() {
			slice = append(slice, s.enumerator.Current())
		}
		if len(slice) < s.count {
			s.skippedEnumerator = &emptyEnumerator[TSource]{}
		} else {
			s.skippedEnumerator = &sliceEnumerator[TSource]{slice: slice[:len(slice)-s.count], cursor: -1}
		}
	}
	return s.skippedEnumerator.MoveNext()
}

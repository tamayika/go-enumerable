package enumerable

func Slice[TSource any](slice []TSource) IEnumerable[TSource] {
	return sliceEnumerable[TSource](slice)
}

type sliceEnumerable[TSource any] []TSource

func (e sliceEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &sliceEnumerator[TSource]{slice: e, cursor: -1}
}

type sliceEnumerator[TSource any] struct {
	slice  []TSource
	cursor int
}

func (s *sliceEnumerator[TSource]) Current() TSource {
	return s.slice[s.cursor]
}

func (s *sliceEnumerator[TSource]) MoveNext() bool {
	s.cursor++
	return s.cursor < len(s.slice)
}

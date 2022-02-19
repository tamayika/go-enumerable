package enumerable

func Empty[TSource any]() IEnumerable[TSource] {
	return emptyEnumerable[TSource]{}
}

type emptyEnumerable[TSource any] struct {
}

func (e emptyEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &emptyEnumerator[TSource]{}
}

type emptyEnumerator[TSource any] struct {
}

func (e *emptyEnumerator[TSource]) Current() TSource {
	panic("empty")
}

func (e *emptyEnumerator[TSource]) MoveNext() bool {
	return false
}

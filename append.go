package enumerable

func Append[TSource any](source IEnumerable[TSource], element TSource) IEnumerable[TSource] {
	return appendEnumerable[TSource]{source, element}
}

type appendEnumerable[TSource any] struct {
	source  IEnumerable[TSource]
	element TSource
}

func (a appendEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &appendEnumerator[TSource]{enumerator: a.source.Enumerator(), element: a.element}
}

type appendEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	element    TSource
	enumrated  bool
}

func (a *appendEnumerator[TSource]) Current() TSource {
	if !a.enumrated {
		return a.enumerator.Current()
	}
	return a.element
}

func (a *appendEnumerator[TSource]) MoveNext() bool {
	if a.enumerator.MoveNext() {
		return true
	}
	if !a.enumrated {
		a.enumrated = true
		return true
	}
	return false
}

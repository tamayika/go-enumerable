package enumerable

func Prepend[TSource any](source IEnumerable[TSource], element TSource) IEnumerable[TSource] {
	return prependEnumerable[TSource]{source, element}
}

type prependEnumerable[TSource any] struct {
	source  IEnumerable[TSource]
	element TSource
}

func (p prependEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &prependEnumerator[TSource]{enumerator: p.source.Enumerator(), element: p.element}
}

type prependEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	element    TSource
	moveCount  int
}

func (p *prependEnumerator[TSource]) Current() TSource {
	switch p.moveCount {
	case 0:
		panic("move next not called")
	case 1:
		return p.element
	}
	return p.enumerator.Current()
}

func (p *prependEnumerator[TSource]) MoveNext() bool {
	p.moveCount++
	if p.moveCount > 1 {
		return p.enumerator.MoveNext()
	}
	return true
}

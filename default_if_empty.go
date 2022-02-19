package enumerable

func DefaultIfEmpty[TSource any](source IEnumerable[TSource]) IEnumerable[TSource] {
	return defaultIfEmptyEnumerable[TSource]{source}
}

type defaultIfEmptyEnumerable[TSource any] struct {
	source IEnumerable[TSource]
}

func (d defaultIfEmptyEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &defaultIfEmptyEnumerator[TSource]{enumerator: d.source.Enumerator(), initial: true}
}

type defaultIfEmptyEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	initial    bool
	empty      bool
}

func (d *defaultIfEmptyEnumerator[TSource]) Current() TSource {
	if d.empty {
		var zero TSource
		return zero
	}
	return d.enumerator.Current()
}

func (d *defaultIfEmptyEnumerator[TSource]) MoveNext() bool {
	if d.initial {
		d.empty = !d.enumerator.MoveNext()
		d.initial = false
		return true
	}
	if d.empty {
		return false
	}
	return d.enumerator.MoveNext()
}

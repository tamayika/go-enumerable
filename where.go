package enumerable

func Where[TSource any](source IEnumerable[TSource], predicate func(TSource) bool) IEnumerable[TSource] {
	return whereEnumerable[TSource]{source, predicate}
}

type whereEnumerable[TSource any] struct {
	source IEnumerable[TSource]
	cb     func(TSource) bool
}

func (w whereEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &whereEnumerator[TSource]{enumerator: w.source.Enumerator(), cb: w.cb}
}

type whereEnumerator[TSource any] struct {
	enumerator IEnumerator[TSource]
	cb         func(TSource) bool
}

func (w *whereEnumerator[TSource]) Current() TSource {
	return w.enumerator.Current()
}

func (w *whereEnumerator[TSource]) MoveNext() bool {
	for w.enumerator.MoveNext() {
		if w.cb(w.enumerator.Current()) {
			return true
		}
	}
	return false
}

package enumerable

func Reverse[TSource any](source IEnumerable[TSource]) IEnumerable[TSource] {
	return reverseEnumerable[TSource]{source}
}

type reverseEnumerable[TSource any] struct {
	source IEnumerable[TSource]
}

func (r reverseEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &reverseEnumerator[TSource]{enumerator: r.source.Enumerator()}
}

type reverseEnumerator[TSource any] struct {
	enumerator         IEnumerator[TSource]
	reversedEnumerator IEnumerator[TSource]
}

func (r *reverseEnumerator[TSource]) Current() TSource {
	return r.reversedEnumerator.Current()
}

func (r *reverseEnumerator[TSource]) MoveNext() bool {
	if r.reversedEnumerator == nil {
		var slice []TSource
		for r.enumerator.MoveNext() {
			slice = append(slice, r.enumerator.Current())
		}
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		r.reversedEnumerator = &sliceEnumerator[TSource]{slice: slice, cursor: -1}
	}
	return r.reversedEnumerator.MoveNext()
}

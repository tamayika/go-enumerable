package enumerable

func Repeat[TResult any](element TResult, count int) IEnumerable[TResult] {
	return repeatEnumerable[TResult]{element, count}
}

type repeatEnumerable[TResult any] struct {
	element TResult
	count   int
}

func (r repeatEnumerable[TResult]) Enumerator() IEnumerator[TResult] {
	return &repeatEnumerator[TResult]{element: r.element, count: r.count}
}

type repeatEnumerator[TResult any] struct {
	element   TResult
	count     int
	moveCount int
}

func (r *repeatEnumerator[TResult]) Current() TResult {
	return r.element
}

func (r *repeatEnumerator[TResult]) MoveNext() bool {
	if r.moveCount < r.count {
		r.moveCount++
		return true
	}
	return false
}

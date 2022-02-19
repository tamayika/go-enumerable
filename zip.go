package enumerable

func Zip[TFirst, TSecond, TResult any](
	first IEnumerable[TFirst], second IEnumerable[TSecond],
	resultSelector func(TFirst, TSecond) TResult,
) IEnumerable[TResult] {
	return zipEnumerable[TFirst, TSecond, TResult]{first, second, resultSelector}
}

type zipEnumerable[TFirst, TSecond, TResult any] struct {
	first          IEnumerable[TFirst]
	second         IEnumerable[TSecond]
	resultSelector func(TFirst, TSecond) TResult
}

func (z zipEnumerable[TFirst, TSecond, TResult]) Enumerator() IEnumerator[TResult] {
	return &zipEnumerator[TFirst, TSecond, TResult]{
		first:          z.first.Enumerator(),
		second:         z.second.Enumerator(),
		resultSelector: z.resultSelector,
	}
}

type zipEnumerator[TFirst, TSecond, TResult any] struct {
	first          IEnumerator[TFirst]
	second         IEnumerator[TSecond]
	resultSelector func(TFirst, TSecond) TResult
}

func (z *zipEnumerator[TFirst, TSecond, TResult]) Current() TResult {
	return z.resultSelector(z.first.Current(), z.second.Current())
}

func (z *zipEnumerator[TFirst, TSecond, TResult]) MoveNext() bool {
	firstOK := z.first.MoveNext()
	secondOK := z.second.MoveNext()
	return firstOK && secondOK
}

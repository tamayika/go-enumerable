package enumerable

func Except[TSource comparable](first, second IEnumerable[TSource]) IEnumerable[TSource] {
	return exceptEnumerable[TSource]{first, second}
}

type exceptEnumerable[TSource comparable] struct {
	first  IEnumerable[TSource]
	second IEnumerable[TSource]
}

func (e exceptEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &exceptEnumerator[TSource]{first: e.first.Enumerator(), second: e.second.Enumerator()}
}

type exceptEnumerator[TSource comparable] struct {
	first        IEnumerator[TSource]
	second       IEnumerator[TSource]
	secondValues map[TSource]struct{}
}

func (e *exceptEnumerator[TSource]) Current() TSource {
	return e.first.Current()
}

func (e *exceptEnumerator[TSource]) MoveNext() bool {
	if e.secondValues == nil {
		e.secondValues = map[TSource]struct{}{}
		for e.second.MoveNext() {
			e.secondValues[e.second.Current()] = struct{}{}
		}
	}
	for e.first.MoveNext() {
		if _, ok := e.secondValues[e.first.Current()]; !ok {
			return true
		}
	}
	return false
}

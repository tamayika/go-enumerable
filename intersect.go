package enumerable

func Intersect[TSource comparable](first, second IEnumerable[TSource]) IEnumerable[TSource] {
	return intersectEnumerable[TSource]{first, second}
}

type intersectEnumerable[TSource comparable] struct {
	first  IEnumerable[TSource]
	second IEnumerable[TSource]
}

func (e intersectEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &intersectEnumerator[TSource]{first: e.first.Enumerator(), second: e.second.Enumerator()}
}

type intersectEnumerator[TSource comparable] struct {
	first        IEnumerator[TSource]
	second       IEnumerator[TSource]
	secondValues map[TSource]struct{}
}

func (e *intersectEnumerator[TSource]) Current() TSource {
	return e.first.Current()
}

func (e *intersectEnumerator[TSource]) MoveNext() bool {
	if e.secondValues == nil {
		e.secondValues = map[TSource]struct{}{}
		for e.second.MoveNext() {
			e.secondValues[e.second.Current()] = struct{}{}
		}
	}
	for e.first.MoveNext() {
		if _, ok := e.secondValues[e.first.Current()]; ok {
			return true
		}
	}
	return false
}

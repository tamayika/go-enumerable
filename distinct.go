package enumerable

func Distinct[TSource comparable](source IEnumerable[TSource]) IEnumerable[TSource] {
	return distinctEnumerable[TSource]{source}
}

type distinctEnumerable[TSource comparable] struct {
	source IEnumerable[TSource]
}

func (d distinctEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &distinctEnumerator[TSource]{enumerator: d.source.Enumerator(), set: map[TSource]struct{}{}}
}

type distinctEnumerator[TSource comparable] struct {
	enumerator IEnumerator[TSource]
	set        map[TSource]struct{}
}

func (d *distinctEnumerator[TSource]) Current() TSource {
	return d.enumerator.Current()
}

func (d *distinctEnumerator[TSource]) MoveNext() bool {
	for d.enumerator.MoveNext() {
		if _, ok := d.set[d.enumerator.Current()]; !ok {
			d.set[d.enumerator.Current()] = struct{}{}
			return true
		}
	}
	return false
}

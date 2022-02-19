package enumerable

import "golang.org/x/exp/constraints"

func Min[TSource constraints.Ordered](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	var min TSource
	var minSet bool
	for enumerator.MoveNext() {
		if !minSet {
			min = enumerator.Current()
			minSet = true
		} else if min > enumerator.Current() {
			min = enumerator.Current()
		}
	}
	if !minSet {
		panic("empty enumerable")
	}
	return min
}

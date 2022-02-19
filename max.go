package enumerable

import "golang.org/x/exp/constraints"

func Max[TSource constraints.Ordered](source IEnumerable[TSource]) TSource {
	enumerator := source.Enumerator()
	var max TSource
	var maxSet bool
	for enumerator.MoveNext() {
		if !maxSet {
			max = enumerator.Current()
			maxSet = true
		} else if max < enumerator.Current() {
			max = enumerator.Current()
		}
	}
	if !maxSet {
		panic("empty enumerable")
	}
	return max
}

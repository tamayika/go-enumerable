package enumerable

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func OrderBy[TSource constraints.Ordered](source IEnumerable[TSource]) IEnumerable[TSource] {
	return orderByEnumerable[TSource]{source}
}

type orderByEnumerable[TSource constraints.Ordered] struct {
	enumarable IEnumerable[TSource]
}

func (o orderByEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &orderByEnumerator[TSource]{enumerator: o.enumarable.Enumerator()}
}

type orderByEnumerator[TSource constraints.Ordered] struct {
	enumerator        IEnumerator[TSource]
	orderedEnumerator IEnumerator[TSource]
}

func (o *orderByEnumerator[TSource]) Current() TSource {
	return o.orderedEnumerator.Current()
}

func (o *orderByEnumerator[TSource]) MoveNext() bool {
	if o.orderedEnumerator == nil {
		var slice []TSource
		for o.enumerator.MoveNext() {
			slice = append(slice, o.enumerator.Current())
		}
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		o.orderedEnumerator = &sliceEnumerator[TSource]{slice: slice, cursor: -1}
	}
	return o.orderedEnumerator.MoveNext()
}

package enumerable

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func OrderByDescending[TSource constraints.Ordered](source IEnumerable[TSource]) IEnumerable[TSource] {
	return orderByDescendingEnumerable[TSource]{source}
}

type orderByDescendingEnumerable[TSource constraints.Ordered] struct {
	enumarable IEnumerable[TSource]
}

func (o orderByDescendingEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &orderByDescendingEnumerator[TSource]{enumerator: o.enumarable.Enumerator()}
}

type orderByDescendingEnumerator[TSource constraints.Ordered] struct {
	enumerator        IEnumerator[TSource]
	orderedEnumerator IEnumerator[TSource]
}

func (o *orderByDescendingEnumerator[TSource]) Current() TSource {
	return o.orderedEnumerator.Current()
}

func (o *orderByDescendingEnumerator[TSource]) MoveNext() bool {
	if o.orderedEnumerator == nil {
		var slice []TSource
		for o.enumerator.MoveNext() {
			slice = append(slice, o.enumerator.Current())
		}
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})
		o.orderedEnumerator = &sliceEnumerator[TSource]{slice: slice, cursor: -1}
	}
	return o.orderedEnumerator.MoveNext()
}

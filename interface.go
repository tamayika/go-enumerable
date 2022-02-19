package enumerable

import "golang.org/x/exp/constraints"

type IEnumerable[TSource any] interface {
	Enumerator() IEnumerator[TSource]
}

type IEnumerator[TSource any] interface {
	Current() TSource
	MoveNext() bool
}

type number interface {
	constraints.Integer | constraints.Float
}

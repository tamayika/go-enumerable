package enumerable

func GroupJoin[TOuter any, TInner any, TKey comparable, TResult any](
	outer IEnumerable[TOuter], inner IEnumerable[TInner],
	outerKeySelector func(TOuter) TKey, innerKeySelector func(TInner) TKey,
	resultSelector func(outer TOuter, innerCollection IEnumerable[TInner]) TResult) IEnumerable[TResult] {
	return groupJoinEnumerable[TOuter, TInner, TKey, TResult]{
		outer,
		inner,
		outerKeySelector,
		innerKeySelector,
		resultSelector,
	}
}

type groupJoinEnumerable[TOuter any, TInner any, TKey comparable, TResult any] struct {
	outer            IEnumerable[TOuter]
	inner            IEnumerable[TInner]
	outerKeySelector func(TOuter) TKey
	innerKeySelector func(TInner) TKey
	resultSelector   func(outer TOuter, innerCollection IEnumerable[TInner]) TResult
}

func (g groupJoinEnumerable[TOuter, TInner, TKey, TResult]) Enumerator() IEnumerator[TResult] {
	return &groupJoinEnumerator[TOuter, TInner, TKey, TResult]{
		outer:            g.outer.Enumerator(),
		inner:            g.inner.Enumerator(),
		outerKeySelector: g.outerKeySelector,
		innerKeySelector: g.innerKeySelector,
		resultSelector:   g.resultSelector,
	}
}

type groupJoinEnumerator[TOuter any, TInner any, TKey comparable, TResult any] struct {
	outer            IEnumerator[TOuter]
	inner            IEnumerator[TInner]
	outerKeySelector func(TOuter) TKey
	innerKeySelector func(TInner) TKey
	resultSelector   func(outer TOuter, innerCollection IEnumerable[TInner]) TResult
	innerMap         map[TKey][]TInner
}

func (g *groupJoinEnumerator[TOuter, TInner, TKey, TResult]) Current() TResult {
	outerKey := g.outerKeySelector(g.outer.Current())
	innerCollection := g.innerMap[outerKey]
	return g.resultSelector(g.outer.Current(), Slice(innerCollection))
}

func (g *groupJoinEnumerator[TOuter, TInner, TKey, TResult]) MoveNext() bool {
	for g.outer.MoveNext() {
		outerKey := g.outerKeySelector(g.outer.Current())
		if g.innerMap == nil {
			g.innerMap = map[TKey][]TInner{}
			for g.inner.MoveNext() {
				innerKey := g.innerKeySelector(g.inner.Current())
				g.innerMap[innerKey] = append(g.innerMap[innerKey], g.inner.Current())
			}
		}
		_, ok := g.innerMap[outerKey]
		if ok {
			return true
		}
	}
	return false
}

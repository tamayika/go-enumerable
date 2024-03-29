package enumerable

func Join[TOuter any, TInner any, TKey comparable, TResult any](
	outer IEnumerable[TOuter], inner IEnumerable[TInner],
	outerKeySelector func(TOuter) TKey, innerKeySelector func(TInner) TKey,
	resultSelector func(TOuter, TInner) TResult) IEnumerable[TResult] {
	return joinEnumerable[TOuter, TInner, TKey, TResult]{
		outer,
		inner,
		outerKeySelector,
		innerKeySelector,
		resultSelector,
	}
}

type joinEnumerable[TOuter any, TInner any, TKey comparable, TResult any] struct {
	outer            IEnumerable[TOuter]
	inner            IEnumerable[TInner]
	outerKeySelector func(TOuter) TKey
	innerKeySelector func(TInner) TKey
	resultSelector   func(TOuter, TInner) TResult
}

func (g joinEnumerable[TOuter, TInner, TKey, TResult]) Enumerator() IEnumerator[TResult] {
	return &joinEnumerator[TOuter, TInner, TKey, TResult]{
		outer:            g.outer.Enumerator(),
		inner:            g.inner.Enumerator(),
		outerKeySelector: g.outerKeySelector,
		innerKeySelector: g.innerKeySelector,
		resultSelector:   g.resultSelector,
	}
}

type joinEnumerator[TOuter any, TInner any, TKey comparable, TResult any] struct {
	outer             IEnumerator[TOuter]
	inner             IEnumerator[TInner]
	outerKeySelector  func(TOuter) TKey
	innerKeySelector  func(TInner) TKey
	resultSelector    func(TOuter, TInner) TResult
	innerMap          map[TKey][]TInner
	currentInner      []TInner
	currentInnerIndex int
}

func (g *joinEnumerator[TOuter, TInner, TKey, TResult]) Current() TResult {
	return g.resultSelector(g.outer.Current(), g.currentInner[g.currentInnerIndex])
}

func (g *joinEnumerator[TOuter, TInner, TKey, TResult]) MoveNext() bool {
	if g.currentInner != nil {
		if g.currentInnerIndex+1 < len(g.currentInner) {
			g.currentInnerIndex++
			return true
		}
		g.currentInner = nil
		g.currentInnerIndex = 0
	}
	for g.outer.MoveNext() {
		outerKey := g.outerKeySelector(g.outer.Current())
		if g.innerMap == nil {
			g.innerMap = map[TKey][]TInner{}
			for g.inner.MoveNext() {
				innerKey := g.innerKeySelector(g.inner.Current())
				g.innerMap[innerKey] = append(g.innerMap[innerKey], g.inner.Current())
			}
		}
		var ok bool
		g.currentInner, ok = g.innerMap[outerKey]
		if ok {
			return true
		}
	}
	return false
}

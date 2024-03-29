package enumerable

func GroupBy[TSource comparable, TKey comparable, TElement any, TResult any](
	source IEnumerable[TSource],
	keySelector func(TSource) TKey, elementSelector func(TSource) TElement,
	resultSelector func(TKey, TElement, TResult) TResult) IEnumerable[TResult] {
	return groupByEnumerable[TSource, TKey, TElement, TResult]{
		source,
		keySelector,
		elementSelector,
		resultSelector,
	}
}

type groupByEnumerable[TSource comparable, TKey comparable, TElement any, TResult any] struct {
	source          IEnumerable[TSource]
	keySelector     func(TSource) TKey
	elementSelector func(TSource) TElement
	resultSelector  func(TKey, TElement, TResult) TResult
}

func (g groupByEnumerable[TSource, TKey, TElement, TResult]) Enumerator() IEnumerator[TResult] {
	return &groupByEnumerator[TSource, TKey, TElement, TResult]{
		enumerator:      g.source.Enumerator(),
		keySelector:     g.keySelector,
		elementSelector: g.elementSelector,
		resultSelector:  g.resultSelector,
	}
}

type groupByEnumerator[TSource, TKey comparable, TElement any, TResult any] struct {
	enumerator       IEnumerator[TSource]
	keySelector      func(TSource) TKey
	elementSelector  func(TSource) TElement
	resultSelector   func(TKey, TElement, TResult) TResult
	resultEnumerator IEnumerator[TResult]
}

func (g *groupByEnumerator[TSource, TKey, TElement, TResult]) Current() TResult {
	return g.resultEnumerator.Current()
}

func (g *groupByEnumerator[TSource, TKey, TElement, TResult]) MoveNext() bool {
	if g.resultEnumerator == nil {
		keys := []TKey{}
		resultMap := map[TKey]TResult{}
		for g.enumerator.MoveNext() {
			key := g.keySelector(g.enumerator.Current())
			element := g.elementSelector(g.enumerator.Current())
			_, ok := resultMap[key]
			resultMap[key] = g.resultSelector(key, element, resultMap[key])
			if !ok {
				keys = append(keys, key)
			}
		}
		results := make([]TResult, len(resultMap))
		for i, key := range keys {
			results[i] = resultMap[key]
		}
		g.resultEnumerator = &sliceEnumerator[TResult]{slice: results, cursor: -1}
	}
	return g.resultEnumerator.MoveNext()
}

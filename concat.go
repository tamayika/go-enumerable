package enumerable

func Concat[TSource any](first, second IEnumerable[TSource]) IEnumerable[TSource] {
	return concatEnumerable[TSource]{first, second}
}

type concatEnumerable[TSource any] struct {
	first  IEnumerable[TSource]
	second IEnumerable[TSource]
}

func (c concatEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &concatEnumerator[TSource]{
		first:  c.first.Enumerator(),
		second: c.second.Enumerator(),
	}
}

type concatEnumerator[TSource any] struct {
	first     IEnumerator[TSource]
	second    IEnumerator[TSource]
	useSecond bool
}

func (c *concatEnumerator[TSource]) Current() TSource {
	if c.useSecond {
		return c.second.Current()
	}
	return c.first.Current()
}

func (c *concatEnumerator[TSource]) MoveNext() bool {
	if !c.useSecond {
		if c.first.MoveNext() {
			return true
		}
		c.useSecond = true
	}
	return c.second.MoveNext()
}

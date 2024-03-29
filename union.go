package enumerable

func Union[TSource comparable](first, second IEnumerable[TSource]) IEnumerable[TSource] {
	return unionEnumerable[TSource]{first, second}
}

type unionEnumerable[TSource comparable] struct {
	first  IEnumerable[TSource]
	second IEnumerable[TSource]
}

func (u unionEnumerable[TSource]) Enumerator() IEnumerator[TSource] {
	return &unionEnumerator[TSource]{
		first:      u.first.Enumerator(),
		second:     u.second.Enumerator(),
		enumerated: map[TSource]struct{}{},
	}
}

type unionEnumerator[TSource comparable] struct {
	first      IEnumerator[TSource]
	second     IEnumerator[TSource]
	enumerated map[TSource]struct{}
	useSecond  bool
}

func (u *unionEnumerator[TSource]) Current() TSource {
	if !u.useSecond {
		return u.first.Current()
	}
	return u.second.Current()
}

func (u *unionEnumerator[TSource]) MoveNext() bool {
	if !u.useSecond {
		for u.first.MoveNext() {
			if _, ok := u.enumerated[u.first.Current()]; !ok {
				u.enumerated[u.first.Current()] = struct{}{}
				return true
			}
		}
		u.useSecond = true
	}
	for u.second.MoveNext() {
		if _, ok := u.enumerated[u.second.Current()]; !ok {
			u.enumerated[u.second.Current()] = struct{}{}
			return true
		}
	}
	return false
}

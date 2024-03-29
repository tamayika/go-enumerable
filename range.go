package enumerable

func Range(start int, count int) IEnumerable[int] {
	return rangeEnumerable{start, count}
}

type rangeEnumerable struct {
	start int
	count int
}

func (r rangeEnumerable) Enumerator() IEnumerator[int] {
	return &rangeEnumerator{current: r.start - 1, end: r.start + r.count - 1}
}

type rangeEnumerator struct {
	current int
	end     int
}

func (r *rangeEnumerator) Current() int {
	return r.current
}

func (r *rangeEnumerator) MoveNext() bool {
	if r.current < r.end {
		r.current++
		return true
	}
	return false
}

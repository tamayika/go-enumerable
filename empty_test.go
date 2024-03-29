package enumerable

import "testing"

func TestEmpty(t *testing.T) {
	enumerator := Empty[int]().Enumerator()
	equal(t, false, enumerator.MoveNext())
	expectPanic(t, func() { enumerator.Current() })
}

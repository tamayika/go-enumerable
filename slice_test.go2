package enumerable

import (
	"testing"
)

func TestSlice(t *testing.T) {
	enumerator := Slice([]string{"a", "b"}).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, "a", enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, "b", enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

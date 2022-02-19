package enumerable

import "testing"

func TestCount(t *testing.T) {
	equal(t, 0, Count(Slice([]string{})))
	equal(t, 1, Count(Slice([]string{"a"})))
	equal(t, 2, Count(Slice([]string{"a", "b"})))
}

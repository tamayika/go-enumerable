package enumerable

import (
	"strconv"
	"testing"
)

func TestToMap(t *testing.T) {
	m := ToMap(Slice([]int{10, 20}), func(v int) int { return v }, func(v int) string { return strconv.Itoa(v) })
	equal(t, 2, len(m))
	equal(t, "10", m[10])
	equal(t, "20", m[20])
}

package enumerable

import "testing"

func TestAverage(t *testing.T) {
	equal(t, 0, Average(Slice([]int{})))
	equal(t, 15, Average(Slice([]int{10, 20})))
	equal(t, 9.5, Average(Slice([]float64{10, 9})))
}

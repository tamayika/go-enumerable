package enumerable

import "testing"

func TestExample(t *testing.T) {
	equal(t, []int{40, 60}, ToSlice(
		Where(
			Select(
				Slice([]int{10, 20, 30}),
				func(v int) int { return v * 2 },
			),
			func(v int) bool { return v > 20 },
		),
	))
}

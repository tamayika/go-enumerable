package enumerable

import "testing"

func TestAggregate(t *testing.T) {
	equal(t, "abc", Aggregate(
		Slice([]string{"a", "b", "c"}),
		"",
		func(ac string, v string) string { return ac + v },
		func(ac string) string { return ac },
	))
}

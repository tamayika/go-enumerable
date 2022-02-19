package enumerable

import "testing"

func TestAppend(t *testing.T){
	enumerator := Append(Slice([]int{10}), 20).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, 10, enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, 20, enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}

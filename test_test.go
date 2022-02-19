package enumerable

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
}

type Pet struct {
	Name  string
	Owner string
}

func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func expectPanic(t *testing.T, cb func()) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("no panic")
		}
	}()
	cb()
}

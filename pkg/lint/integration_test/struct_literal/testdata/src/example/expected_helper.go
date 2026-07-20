package example

import (
	"example/inner"
	"testing"
)

func ExpectedHelper(t *testing.T) {
	assertStruct(t, &inner.MyStruct{})
	buildStruct(t, &inner.MyStruct{})
}

func assertStruct(t *testing.T, expected *inner.MyStruct) {
	t.Helper()
	_ = expected
}

func buildStruct(t *testing.T, value *inner.MyStruct) {
	t.Helper()
	_ = value
}

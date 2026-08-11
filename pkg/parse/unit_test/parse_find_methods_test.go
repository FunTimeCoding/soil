package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

func chainedSource() string {
	return "package test\n\nfunc Run() {\n\tlayout.New(i).WithTheme(theme.Straw).WithCommandPalette(\"/palette\")\n}\n"
}

func routeSource() string {
	return "package test\n\nfunc Run() {\n\tm.HandleFunc(\"GET /palette\", a)\n\tm.HandleFunc(\"GET /favicon.ico\", b)\n}\n"
}

func TestFindMethodsChained(t *testing.T) {
	f, _, e := parse.Source("test.go", chainedSource())
	assert.Nil(t, e)
	assert.Integer(t, 1, len(parse.FindMethods(f, "WithTheme")))
	assert.Integer(t, 1, len(parse.FindMethods(f, "WithCommandPalette")))
}

func TestFindMethodsAbsent(t *testing.T) {
	f, _, e := parse.Source("test.go", chainedSource())
	assert.Nil(t, e)
	assert.Integer(t, 0, len(parse.FindMethods(f, "WithBrandNode")))
}

func TestFindMethodsRepeated(t *testing.T) {
	f, _, e := parse.Source("test.go", routeSource())
	assert.Nil(t, e)
	assert.Integer(t, 2, len(parse.FindMethods(f, "HandleFunc")))
}

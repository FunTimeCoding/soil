package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

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

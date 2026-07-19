package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

func TestSource(t *testing.T) {
	f, s, e := parse.Source(
		"test.go",
		"package test\n\nvar Identity = \"hello\"\n",
	)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	assert.NotNil(t, s)
	assert.String(t, "test", f.Name.Name)
}

func TestSourceInvalid(t *testing.T) {
	_, _, e := parse.Source("test.go", "not valid go")
	assert.NotNil(t, e)
}

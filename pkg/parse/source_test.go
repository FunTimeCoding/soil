package parse

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSource(t *testing.T) {
	f, s, e := Source(
		"test.go",
		"package test\n\nvar Identity = \"hello\"\n",
	)
	assert.Nil(t, e)
	assert.NotNil(t, f)
	assert.NotNil(t, s)
	assert.String(t, "test", f.Name.Name)
}

func TestSourceInvalid(t *testing.T) {
	_, _, e := Source("test.go", "not valid go")
	assert.NotNil(t, e)
}

package parse

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestImportNameUnaliased(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n",
	)
	assert.Nil(t, e)
	name, found := ImportName(f, "github.com/example/reporter")
	assert.True(t, found)
	assert.String(t, "reporter", name)
}

func TestImportNameAliased(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport r \"github.com/example/reporter\"\n",
	)
	assert.Nil(t, e)
	name, found := ImportName(f, "github.com/example/reporter")
	assert.True(t, found)
	assert.String(t, "r", name)
}

func TestImportNameNotFound(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport \"fmt\"\n",
	)
	assert.Nil(t, e)
	_, found := ImportName(f, "github.com/example/reporter")
	assert.False(t, found)
}

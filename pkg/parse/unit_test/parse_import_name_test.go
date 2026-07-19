package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

func TestImportNameUnaliased(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n",
	)
	assert.Nil(t, e)
	name, found := parse.ImportName(f, "github.com/example/reporter")
	assert.True(t, found)
	assert.String(t, "reporter", name)
}

func TestImportNameAliased(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport r \"github.com/example/reporter\"\n",
	)
	assert.Nil(t, e)
	name, found := parse.ImportName(f, "github.com/example/reporter")
	assert.True(t, found)
	assert.String(t, "r", name)
}

func TestImportNameNotFound(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport \"fmt\"\n",
	)
	assert.Nil(t, e)
	_, found := parse.ImportName(f, "github.com/example/reporter")
	assert.False(t, found)
}

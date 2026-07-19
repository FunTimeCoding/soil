package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

func TestHasCallFound(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n\nfunc Run() {\n\treporter.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.True(t, parse.HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallNotImported(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nfunc Run() {\n\treporter.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.False(t, parse.HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallAliased(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport r \"github.com/example/reporter\"\n\nfunc Run() {\n\tr.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.True(t, parse.HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallWrongFunction(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n\nfunc Run() {\n\treporter.Start(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.False(t, parse.HasCall(f, "github.com/example/reporter", "New"))
}

package parse

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestHasCallFound(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n\nfunc Run() {\n\treporter.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.True(t, HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallNotImported(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nfunc Run() {\n\treporter.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.False(t, HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallAliased(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport r \"github.com/example/reporter\"\n\nfunc Run() {\n\tr.New(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.True(t, HasCall(f, "github.com/example/reporter", "New"))
}

func TestHasCallWrongFunction(t *testing.T) {
	f, _, e := Source(
		"test.go",
		"package test\n\nimport \"github.com/example/reporter\"\n\nfunc Run() {\n\treporter.Start(\"test\")\n}\n",
	)
	assert.Nil(t, e)
	assert.False(t, HasCall(f, "github.com/example/reporter", "New"))
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/parse"
	"testing"
)

func TestFindCallsCollectsAll(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nfunc register() {\n\tadd(mcp.NewTool(\"alfa\"))\n\tadd(mcp.NewTool(\"bravo\"))\n}\n",
	)
	assert.Nil(t, e)
	result := parse.FindCalls(f, "mcp", "NewTool")
	assert.Integer(t, 2, len(result))
}

func TestFindCallsNested(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nfunc Tool() any {\n\treturn mcp.NewTool(\"alfa\", mcp.WithString(\"name\"))\n}\n",
	)
	assert.Nil(t, e)
	result := parse.FindCalls(f, "mcp", "NewTool")
	assert.Integer(t, 1, len(result))
}

func TestFindCallsNone(t *testing.T) {
	f, _, e := parse.Source(
		"test.go",
		"package test\n\nfunc Run() {\n\tother.NewTool(\"alfa\")\n}\n",
	)
	assert.Nil(t, e)
	assert.Integer(t, 0, len(parse.FindCalls(f, "mcp", "NewTool")))
}

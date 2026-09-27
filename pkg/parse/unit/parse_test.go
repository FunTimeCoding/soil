package unit

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
	f, _, e := parse.Source("test.go", "package test\n\nimport \"fmt\"\n")
	assert.Nil(t, e)
	_, found := parse.ImportName(f, "github.com/example/reporter")
	assert.False(t, found)
}

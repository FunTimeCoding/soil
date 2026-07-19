package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func TestAddConcernRelativizesPath(t *testing.T) {
	r := output.NewResultsWithDirectory("/Users/example/src/soil/")
	r.AddConcern(
		concern.NewFile(
			"test",
			"finding",
			"/Users/example/src/soil/pkg/foo/bar.go",
			false,
		),
	)
	assert.Integer(t, 1, len(r.Entries))
	assert.String(t, "pkg/foo/bar.go", r.Entries[0].Path)
}

func TestAddConcernPreservesFields(t *testing.T) {
	r := output.NewResults()
	r.AddConcern(
		concern.NewLine(
			constant.BlankInsideFunctionKey,
			"Blank line",
			"pkg/foo.go",
			5,
			"original line",
			true,
		),
	)
	assert.Integer(t, 1, len(r.Entries))
	assert.String(t, constant.BlankInsideFunctionKey, r.Entries[0].Key)
	assert.String(t, "Blank line", r.Entries[0].Text)
	assert.String(t, concern.Line, r.Entries[0].Type)
	assert.Integer(t, 5, r.Entries[0].Line)
	assert.Boolean(t, true, r.Entries[0].Fixed)
}

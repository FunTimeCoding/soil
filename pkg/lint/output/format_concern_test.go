package output

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"testing"
)

func TestFormatLineConcern(t *testing.T) {
	assert.String(
		t,
		"pkg/foo/bar.go:5: Blank line between statements",
		formatConcern(
			concern.NewLine(
				constant.BlankInsideFunctionKey,
				"Blank line between statements",
				"pkg/foo/bar.go",
				5,
				"",
				false,
			),
		),
	)
}

func TestFormatLineConcernApplied(t *testing.T) {
	assert.String(
		t,
		"pkg/foo/bar.go:5: Blank line between statements (auto-fixed)",
		formatConcern(
			concern.NewLine(
				constant.BlankInsideFunctionKey,
				"Blank line between statements",
				"pkg/foo/bar.go",
				5,
				"",
				true,
			),
		),
	)
}

func TestFormatFileConcern(t *testing.T) {
	assert.String(
		t,
		"pkg/tool/goaudit/run.go: multiple identities in one file: Run and runTable",
		formatConcern(
			concern.NewFile(
				"file_identity",
				"multiple identities in one file: Run and runTable",
				"pkg/tool/goaudit/run.go",
				false,
			),
		),
	)
}

func TestFormatFileConcernApplied(t *testing.T) {
	assert.String(
		t,
		"pkg/types/dir_path.go: renamed DirPath → DirectoryPath (2 references) (auto-fixed)",
		formatConcern(
			concern.NewFile(
				"renamed",
				"renamed DirPath → DirectoryPath (2 references)",
				"pkg/types/dir_path.go",
				true,
			),
		),
	)
}

func TestFormatPackageConcern(t *testing.T) {
	assert.String(
		t,
		"pkg/tool/godarwin: no standard service suffix (expected *d)",
		formatConcern(
			concern.NewPackage(
				"missing_suffix",
				"no standard service suffix (expected *d)",
				"pkg/tool/godarwin",
			),
		),
	)
}

func TestFormatLineZeroNoLineNumber(t *testing.T) {
	assert.String(
		t,
		"pkg/foo.go: finding",
		formatConcern(
			concern.NewLine(
				"test",
				"finding",
				"pkg/foo.go",
				0,
				"",
				false,
			),
		),
	)
}

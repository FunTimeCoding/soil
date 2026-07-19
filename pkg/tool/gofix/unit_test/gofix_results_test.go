package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func TestRelativize(t *testing.T) {
	t.Run(
		"StripsPrefix",
		func(t *testing.T) {
			assert.String(
				t,
				"pkg/foo/bar.go",
				output.NewResultsWithDirectory(
					"/Users/example/src/go-example/",
				).Relativize(
					"/Users/example/src/go-example/pkg/foo/bar.go",
				),
			)
		},
	)
	t.Run(
		"NonMatchingPassesThrough",
		func(t *testing.T) {
			assert.String(
				t,
				"/other/path/file.go",
				output.NewResultsWithDirectory(
					"/Users/example/src/go-example/",
				).Relativize("/other/path/file.go"),
			)
		},
	)
	t.Run(
		"EmptyPrefixPassesThrough",
		func(t *testing.T) {
			assert.String(
				t,
				"/some/path/file.go",
				output.NewResultsWithDirectory(
					"",
				).Relativize("/some/path/file.go"),
			)
		},
	)
}

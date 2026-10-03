package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"testing"
)

func assertSecondRunQuiet(
	t *testing.T,
	directory string,
) {
	t.Helper()
	before := readGoFiles(t, directory)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunFormatFixWithDirectory([]string{"./..."}, directory, false, r)
	assert.Count(t, 0, filterApplied(r.Entries))
	assert.Any(t, before, readGoFiles(t, directory))
}

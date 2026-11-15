package fix_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"testing"
)

func SecondRunQuiet(
	t *testing.T,
	directory string,
) {
	t.Helper()
	before := GoFiles(t, directory)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunFormatFix([]string{goModConstant.AllPackages}, directory, false, r)
	assert.Count(t, 0, Applied(r.Entries))
	assert.Any(t, before, GoFiles(t, directory))
}

package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"testing"
)

func checkModule(
	t *testing.T,
	source string,
) *output.Results {
	t.Helper()
	directory := testutil.PrepareTestPackage(t, source)
	p, results := testutil.LoadFromDirectory(t, directory)
	all, _, e := resolve.LoadPackages(directory, "./...")
	assert.FatalOnError(t, e)
	unclosed_resource.Check(p, results, unclosed_resource.NewSummaries(all))

	return results
}

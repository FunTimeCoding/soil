package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func (o *Tester) UnclosedResource(source string) *output.Results {
	o.t.Helper()
	directory := testutil.PrepareTestPackage(o.t, source)
	p, results := testutil.LoadDirectory(o.t, directory)
	all, _, e := resolve.LoadPackages(directory, goModConstant.AllPackages)
	assert.FatalOnError(o.t, e)
	unclosed_resource.Check(p, results, unclosed_resource.NewSummaries(all))

	return results
}

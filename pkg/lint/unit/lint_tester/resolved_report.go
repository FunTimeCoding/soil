package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
)

func (o *Tester) ResolvedReport(
	root string,
	f *warm_forest.Forest,
	c *pointer_cache.Cache,
) *output.Results {
	o.t.Helper()
	c.Begin(f.Refresh())
	p := option.New("", false)
	p.Trees = f
	p.Resolved = c

	return lint.Report("golint", root, p)
}

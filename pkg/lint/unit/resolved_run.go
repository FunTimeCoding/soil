package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"testing"
)

func resolvedRun(
	t *testing.T,
	root string,
	f *warm_forest.Forest,
	c *pointer_cache.Cache,
) *output.Results {
	t.Helper()
	c.Begin(f.Refresh())
	o := option.New("", false)
	o.Trees = f
	o.Resolved = c

	return lint.Report("golint", root, o)
}

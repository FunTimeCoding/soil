package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func SameRewrite(
	t *testing.T,
	run func(
		s *service_tester.Services,
		directory string,
	) (*output.Results, error),
) string {
	t.Helper()
	var trees []string

	for _, full := range []bool{false, true} {
		directory := service_tester.Prepare(t, "indexed/src")
		s := service_tester.NewService(t)

		if full {
			s.Workspace.UseFullLoad()
		}

		r, e := run(s, directory)
		assert.FatalOnError(t, e)
		testutil.AssertBlocked(t, r, 0)
		trees = append(trees, DirectoryTree(t, directory))
	}

	assert.String(t, trees[1], trees[0])

	return trees[0]
}

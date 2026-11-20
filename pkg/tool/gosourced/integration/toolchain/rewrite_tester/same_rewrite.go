package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) SameRewrite(
	run func(
		s *service_tester.Services,
		directory string,
	) (*output.Results, error),
) string {
	o.t.Helper()
	var trees []string

	for _, full := range []bool{false, true} {
		s := service_tester.New(o.t, service_tester.WithFixture("indexed/src"))

		if full {
			s.Workspace.UseFullLoad()
		}

		r, e := run(s.Services, s.Directory)
		assert.FatalOnError(o.t, e)
		testutil.AssertBlocked(o.t, r, 0)
		trees = append(trees, o.DirectoryTree(s.Directory))
	}

	assert.String(o.t, trees[1], trees[0])

	return trees[0]
}

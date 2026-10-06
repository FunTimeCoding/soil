package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func sameRewrite(
	t *testing.T,
	run func(s *service.Service, directory string) (*output.Results, error),
) string {
	t.Helper()
	var trees []string

	for _, full := range []bool{false, true} {
		directory := testutil.PrepareTestPackage(
			t,
			service_tester.ServiceTestdata("indexed/src"),
		)
		s := testService()

		if full {
			s.UseFullLoad()
		}

		r, e := run(s, directory)
		assert.FatalOnError(t, e)
		testutil.AssertBlocked(t, r, 0)
		trees = append(trees, treeOf(t, directory))
	}

	assert.String(t, trees[1], trees[0])

	return trees[0]
}

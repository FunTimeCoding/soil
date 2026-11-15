package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func RemovalRefused(
	t *testing.T,
	expect string,
	function string,
) {
	t.Helper()
	d := service_tester.Prepare(t, "remove-parameters-refused/src")
	before := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				function,
				"",
				[]string{"version"},
			),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, expect)
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go"),
	)
}

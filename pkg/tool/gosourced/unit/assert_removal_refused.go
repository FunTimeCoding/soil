package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func assertRemovalRefused(
	t *testing.T,
	expect string,
	function string,
) {
	t.Helper()
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters-refused/src"),
	)
	before := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	r, e := testService().RemoveParameters(
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

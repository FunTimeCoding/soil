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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters-refused/src"),
	)
	before := o.Read("pkg/caller/caller.go")
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
	assert.String(t, before, o.Read("pkg/caller/caller.go"))
}

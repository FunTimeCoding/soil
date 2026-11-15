package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func AdditionRefused(
	t *testing.T,
	expect string,
	a *addition.Parameter,
) {
	t.Helper()
	d := service_tester.Prepare(t, "add-parameters/src")
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go")
	target := service_tester.ReadFixtureFile(t, d, "pkg/target/target.go")
	r, e := service_tester.NewService(t).Symbol.AddParameters(
		t.Context(),
		d,
		[]*addition.Parameter{a},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, expect)
	assert.String(
		t,
		caller,
		service_tester.ReadFixtureFile(t, d, "pkg/caller/caller.go"),
	)
	assert.String(
		t,
		target,
		service_tester.ReadFixtureFile(t, d, "pkg/target/target.go"),
	)
}

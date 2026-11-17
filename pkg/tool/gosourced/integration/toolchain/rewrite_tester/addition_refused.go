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
	o := service_tester.New(t, service_tester.WithFixture("add-parameters/src"))
	caller := o.Read("pkg/caller/caller.go")
	target := o.Read("pkg/target/target.go")
	r, e := o.Symbol.AddParameters(
		t.Context(),
		o.Directory,
		[]*addition.Parameter{a},
		nil,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, expect)
	assert.String(t, caller, o.Read("pkg/caller/caller.go"))
	assert.String(t, target, o.Read("pkg/target/target.go"))
}

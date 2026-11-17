package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func NamingRefused(
	t *testing.T,
	expect string,
	function string,
) {
	t.Helper()
	o := service_tester.New(
		t,
		service_tester.WithFixture("name-parameters-refused/src"),
	)
	before := o.Read("pkg/handler/handler.go")
	r, e := o.Symbol.NameParameters(
		t.Context(),
		o.Directory,
		naming.NewRequest(
			"example/pkg/handler",
			[]*naming.Function{naming.NewFunction(function, "")},
			"context.Context",
			"x",
			nil,
		),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, expect)
	assert.String(t, before, o.Read("pkg/handler/handler.go"))
}

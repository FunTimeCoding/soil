package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func assertNamingRefused(
	t *testing.T,
	expect string,
	function string,
) {
	t.Helper()
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("name-parameters-refused/src"),
	)
	before := service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go")
	r, e := service_tester.NewService(t).Symbol.NameParameters(
		t.Context(),
		d,
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
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/handler/handler.go"),
	)
}

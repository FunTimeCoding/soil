package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) NamingRefused(
	expect string,
	function string,
) {
	o.t.Helper()
	s := service_tester.New(
		o.t,
		service_tester.WithFixture("name-parameters-refused/src"),
	)
	before := s.Read("pkg/handler/handler.go")
	r, e := s.Symbol.NameParameters(
		o.t.Context(),
		s.Directory,
		naming.NewRequest(
			"example/pkg/handler",
			[]*naming.Function{naming.NewFunction(function, "")},
			"context.Context",
			"x",
			nil,
		),
		false,
	)
	assert.FatalOnError(o.t, e)
	testutil.AssertBlockedContains(o.t, r, expect)
	assert.String(o.t, before, s.Read("pkg/handler/handler.go"))
}

package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) AdditionRefused(
	expect string,
	a *addition.Parameter,
) {
	o.t.Helper()
	s := service_tester.New(
		o.t,
		service_tester.WithFixture("add-parameters/src"),
	)
	caller := s.Read("pkg/caller/caller.go")
	target := s.Read("pkg/target/target.go")
	r, e := s.Symbol.AddParameters(
		o.t.Context(),
		s.Directory,
		[]*addition.Parameter{a},
		nil,
		false,
	)
	assert.FatalOnError(o.t, e)
	testutil.AssertBlockedContains(o.t, r, expect)
	assert.String(o.t, caller, s.Read("pkg/caller/caller.go"))
	assert.String(o.t, target, s.Read("pkg/target/target.go"))
}

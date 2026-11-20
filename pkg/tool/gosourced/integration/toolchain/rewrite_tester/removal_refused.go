package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) RemovalRefused(
	expect string,
	function string,
) {
	o.t.Helper()
	s := service_tester.New(
		o.t,
		service_tester.WithFixture("remove-parameters-refused/src"),
	)
	before := s.Read("pkg/caller/caller.go")
	r, e := s.Symbol.RemoveParameters(
		o.t.Context(),
		s.Directory,
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
	assert.FatalOnError(o.t, e)
	testutil.AssertBlockedContains(o.t, r, expect)
	assert.String(o.t, before, s.Read("pkg/caller/caller.go"))
}

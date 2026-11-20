package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) MoveMethods(
	methods ...string,
) (*service_tester.Tester, *output.Results) {
	o.t.Helper()
	s := service_tester.New(o.t, service_tester.WithFixture("move-methods/src"))
	r, e := s.Symbol.MoveMethods(
		o.t.Context(),
		s.Directory,
		"example/pkg/alfa",
		"Service",
		methods,
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(o.t, e)

	return s, r
}

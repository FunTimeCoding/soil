package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func (o *Tester) MoveHeldGreet(s *service_tester.Tester) *output.Results {
	o.t.Helper()
	r, e := s.Symbol.MoveMethods(
		o.t.Context(),
		s.Directory,
		"example/pkg/alfa",
		"Service",
		[]string{"Greet"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(o.t, e)

	return r
}

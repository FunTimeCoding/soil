package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func MoveHeldGreet(
	t *testing.T,
	o *service_tester.Tester,
) *output.Results {
	t.Helper()
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Service",
		[]string{"Greet"},
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)

	return r
}

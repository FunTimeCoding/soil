package module_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func MoveMethods(
	t *testing.T,
	methods ...string,
) (*service_tester.Tester, *output.Results) {
	t.Helper()
	o := service_tester.New(t, service_tester.WithFixture("move-methods/src"))
	r, e := o.Symbol.MoveMethods(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"Service",
		methods,
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)

	return o, r
}

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
) (string, *output.Results) {
	t.Helper()
	d := service_tester.Prepare(t, "move-methods/src")
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
		"example/pkg/alfa",
		"Service",
		methods,
		"example/pkg/bravo",
		"Worker",
		"",
		false,
	)
	assert.FatalOnError(t, e)

	return d, r
}

package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func moveFixtureMethods(
	t *testing.T,
	methods ...string,
) (string, *output.Results) {
	t.Helper()
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("move-methods/src"),
	)
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

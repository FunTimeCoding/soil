package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func moveHeldGreet(
	t *testing.T,
	d string,
) *output.Results {
	t.Helper()
	r, e := service_tester.NewService(t).Symbol.MoveMethods(
		t.Context(),
		d,
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

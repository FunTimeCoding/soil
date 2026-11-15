package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"testing"
)

func Reports(
	t *testing.T,
	output string,
) *model_context_server.Server {
	t.Helper()

	return Mount(
		t,
		service_tester.NewOption(
			t,
			t.TempDir(),
			output,
			service_tester.NewParser(t),
		),
		memory.New(),
	)
}

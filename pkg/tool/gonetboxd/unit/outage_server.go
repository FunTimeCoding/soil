package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func outageServer(
	t *testing.T,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()

	return newNetboxServer(t, c, r)
}

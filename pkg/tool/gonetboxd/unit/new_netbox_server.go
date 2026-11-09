package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"testing"
)

func newNetboxServer(
	t *testing.T,
	c face.NetBoxClient,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()

	return newDaemonServer(t, newDaemon(t, c), r)
}

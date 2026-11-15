package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"testing"
)

func New(
	t *testing.T,
	c face.NetBoxClient,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()

	return Mount(t, Daemon(t, c), r)
}

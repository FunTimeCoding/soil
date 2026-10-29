package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/service"
	"testing"
)

func mountedUpload(
	t *testing.T,
	v *service.Service,
) *model_context_server.Server {
	t.Helper()

	return mountedReported(t, v, memory.New())
}

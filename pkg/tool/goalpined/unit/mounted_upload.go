package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"testing"
)

func mountedUpload(
	t *testing.T,
	o *option.Alpine,
) *model_context_server.Server {
	t.Helper()

	return mountedReported(t, o, memory.New())
}

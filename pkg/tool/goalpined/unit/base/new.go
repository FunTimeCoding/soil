package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"testing"
)

func New(
	t *testing.T,
	o *option.Alpine,
) *model_context_server.Server {
	t.Helper()

	return Mount(t, o, memory.New())
}

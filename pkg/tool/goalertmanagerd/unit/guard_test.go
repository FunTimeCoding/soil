package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			healthyDaemon(t).Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}

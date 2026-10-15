package unit

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/provision"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func newTriggerServer(
	t *testing.T,
	n *mock_runner.Runner,
	r face.Reporter,
) *model_context_server.Server {
	t.Helper()
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			provision.Mount(n, r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}

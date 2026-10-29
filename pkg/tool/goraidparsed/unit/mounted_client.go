package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func mountedClient(
	t *testing.T,
	g face.Generator,
) *client.Client {
	t.Helper()
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			m *guard.Mux,
		) {
			goraidparsed.Mount(g, memory.New(), mock_recorder.New(), m)
		},
	)
	t.Cleanup(v.Close)
	result := client.New(
		locator.New("localhost").Port(v.Port).Insecure(),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}

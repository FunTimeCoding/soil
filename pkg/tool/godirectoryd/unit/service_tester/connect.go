package service_tester

import (
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func Connect(
	t *testing.T,
	c *mock_client.Client,
	m *memory.Memory,
) *model_context_client.Client {
	t.Helper()
	v := Shared(t, c, m)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			godirectoryd.Mount(v, web.New(v), m, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := s.Connect(t)
	t.Cleanup(result.Close)

	return result
}

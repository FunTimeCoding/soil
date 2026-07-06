package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			model_context.New(
				c,
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)

	return &Server{MockClient: c, ContextServer: v}
}

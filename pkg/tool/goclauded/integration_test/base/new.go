package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	generated "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := service_tester.New(t)
	l := logger.New(t.Context())

	return &Server{
		Tester: s,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(
							s.Service,
							l,
							memory.New(),
							t.TempDir(),
							t.TempDir(),
						),
						nil,
					),
					m,
				)
				model_context.New(
					s.Service,
					memory.New(),
					l,
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(m)
				web.New(s.Service).Mount(m)
			},
		),
	}
}

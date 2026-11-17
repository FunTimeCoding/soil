package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func Serve(
	t *testing.T,
	s *service_tester.Tester,
) *model_context_server.Server {
	t.Helper()
	d := s.Daemon()

	return model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(s.Reporter, mock_recorder.New(), g)
		},
	)
}

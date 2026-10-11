package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosourced"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	fixture string,
) *Server {
	t.Helper()
	d := testutil.PrepareTestPackage(t, fixture)
	i := inventory.New()
	i.Modules = []module.Module{{Name: "test", Directory: d}}
	s := service.New(i)
	s.UseIndex(t.TempDir())
	v := model_context_server.New(
		t,
		func(
			m *http.ServeMux,
			g *guard.Mux,
		) {
			gosourced.Mount(s, memory.New(), mock_recorder.New(), g)
		},
	)
	result := &Server{Server: v, Directory: d}
	t.Cleanup(result.Close)

	return result
}

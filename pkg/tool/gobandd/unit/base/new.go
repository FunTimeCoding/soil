package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	d, engines := daemon_tester.New(t, "alfa", "bravo")
	alfa := engines["alfa"]
	bravo := engines["bravo"]
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	c, e := client.NewClientWithResponses(
		v.Address(webConstant.RootPath),
		client.WithHTTPClient(v.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generative.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	result := &Server{Alfa: alfa, Bravo: bravo, Client: c, Server: v}
	t.Cleanup(result.Close)

	return result
}

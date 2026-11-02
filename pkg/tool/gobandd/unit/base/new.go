package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gobandd"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	alfa := mock_client.New()
	bravo := mock_client.New()
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example"},
		instance.Instance{Name: "bravo", Host: "bravo.example"},
	)
	s := service.New(o, logger.New(t.Context()), memory.New())
	s.Pool().Set("alfa", alfa)
	s.Pool().Set("bravo", bravo)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gobandd.Mount(s, memory.New(), mock_recorder.New(), g)
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

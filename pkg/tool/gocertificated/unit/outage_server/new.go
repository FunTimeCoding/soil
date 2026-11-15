package outage_server

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	f := service_tester.NewForge(t)
	f.Fail(outage_tester.RefusedDial())
	d := daemon_tester.New(t, store_tester.NewOption(t, f), r)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
		model_context_server.WithSignOn(),
	)
	t.Cleanup(result.Close)
	c, e := client.NewClientWithResponses(
		result.Address(""),
		client.WithHTTPClient(result.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	_, h := c.PostAuthorityWithBodyWithResponse(
		t.Context(),
		webConstant.Object,
		bytes.NewReader(notation.Marshal(server_tester.RootBody())),
	)
	assert.FatalOnError(t, h)

	return result, outage_tester.Guard(
		t,
		d.Workers(),
		constant.ForgePeer,
	).Unavailable
}

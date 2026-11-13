package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func outageServer(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	f := service_tester.NewForge(t)
	f.Fail(refusedConnection())
	d := daemon_tester.New(t, store_tester.NewOption(t, f), r)
	result := model_context_server.NewSignedOn(
		t,
		authorization_tester.Client,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
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
	postAuthority(t, c, server_tester.RootBody())

	return result, d.Workers()[0].(*peer.Guard).Unavailable
}

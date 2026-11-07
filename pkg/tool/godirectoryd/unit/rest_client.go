package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func restClient(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *client.ClientWithResponses {
	t.Helper()
	d := daemon_tester.New(t, c, r)
	s := model_context_server.NewSignedOn(
		t,
		authorization_tester.Client,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result
}

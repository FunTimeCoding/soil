package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func restClient(
	t *testing.T,
	c face.JellyfinSource,
	r *memory.Memory,
) *client.ClientWithResponses {
	t.Helper()
	d := daemon_tester.New(t, c)
	s := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(constant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result
}

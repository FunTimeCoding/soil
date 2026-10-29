package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (t *Tester) Mounted(
	readOnly bool,
) (*model_context_client.Client, *memory.Memory) {
	t.t.Helper()
	r := memory.New()
	v := model_context_server.New(
		t.t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gokubernetesd.Mount(t.Service, readOnly, r, mock_recorder.New(), g)
		},
	)
	t.t.Cleanup(v.Close)
	c := v.Connect(t.t)
	t.t.Cleanup(c.Close)
	c.MustCallTool(
		constant.UseCluster,
		map[string]any{"cluster": t.Cluster.Name()},
	)

	return c, r
}

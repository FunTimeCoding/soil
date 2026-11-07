package terraform_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/daemon"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func NewServedWithDownstream(
	t *testing.T,
	c *mock_client.Client,
	plan string,
	downstream *provisionMock.Client,
) *Tester {
	t.Helper()
	result := fixture(t, c, constant.FixtureHostname, 0, plan, downstream)
	d := daemon.New(
		result.Option,
		logger.New(t.Context()),
		result.Reporter,
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)

	for _, w := range d.Workers() {
		w.Start()
	}

	result.served = model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(result.Reporter, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.served.Close)
	result.waitForApply()

	return result
}

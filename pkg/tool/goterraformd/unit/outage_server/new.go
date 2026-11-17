package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	downstream := provisionMock.New(t, "alfa.example")
	downstream.Fail(downstream.Reachable, outage_tester.RefusedDial())
	downstream.Fail(downstream.PostTrigger, outage_tester.RefusedDial())
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "runs.db")
	o.Cluster = func() face.KubernetesClient { return mock_client.New() }
	o.DownstreamClients = func() []provisionFace.Client {
		return []provisionFace.Client{downstream}
	}
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	workers := d.Workers()
	workers[len(workers)-1].Close()
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)

	return result, outage_tester.Guard(
		t,
		workers,
		join.Colon(constant.DownstreamPeer, downstream.Address()),
	).Unavailable
}

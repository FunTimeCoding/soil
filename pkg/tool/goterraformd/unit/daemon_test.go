package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	kubernetesMock "github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"path/filepath"
	"testing"
)

func TestDaemonHandsOutItsGuardsFirst(t *testing.T) {
	workers := base.New(
		t,
		kubernetesMock.New(),
		service_tester.WithDownstream(
			mock_client.New(t, "alfa:8080"),
			mock_client.New(t, "bravo:8080"),
		),
	).Daemon().Workers()
	assert.Count(t, 4, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestDaemonCountsReapsOnItsRegistry(t *testing.T) {
	s := base.New(t, kubernetesMock.New())
	s.Daemon()
	assert.True(t, reaper.CountsOn(s.Registry))
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "goterraformd.db")
	o.Cluster = func() face.KubernetesClient { return kubernetesMock.New() }
	o.DownstreamClients = func() []provisionFace.Client {
		return []provisionFace.Client{mock_client.New(t, "alfa:8080")}
	}
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

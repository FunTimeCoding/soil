package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsTwoGuardsPerInstanceToLifecycle(t *testing.T) {
	assert.Count(t, 4, healthyDaemon(t, "alfa", "bravo").Workers())
}

func TestClosingDaemonClosesBothClientsOfEveryInstance(t *testing.T) {
	a := mock_client.New()
	p := prometheusMock.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": a},
		map[string]*prometheusMock.Client{"alfa": p},
	)
	d.Close()
	assert.Integer(t, 1, a.Closed)
	assert.Integer(t, 1, p.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	alertmanagers := map[string]*mock_client.Client{"alfa": mock_client.New()}
	prometheuses := map[string]*prometheusMock.Client{
		"alfa": prometheusMock.New(),
	}
	before := goleak.IgnoreCurrent()
	d := daemon_tester.New(t, alertmanagers, prometheuses)
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

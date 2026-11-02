package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonHandsOneGuardPerInstanceToTheLifecycle(t *testing.T) {
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{
			"alfa":  mock_client.New(),
			"bravo": mock_client.New(),
		},
	)
	assert.Count(t, 2, d.Workers())
}

func TestClosingTheDaemonClosesEveryInstanceClient(t *testing.T) {
	alfa := mock_client.New()
	bravo := mock_client.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": bravo},
	)
	d.Close()
	assert.Integer(t, 1, alfa.Closed)
	assert.Integer(t, 1, bravo.Closed)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	clients := map[string]*mock_client.Client{
		"alfa":  mock_client.New(),
		"bravo": mock_client.New(),
	}
	before := goleak.IgnoreCurrent()
	d := daemon_tester.New(t, clients)
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

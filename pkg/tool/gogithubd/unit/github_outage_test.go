package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/github/mock_client"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
	"testing"
)

func TestAGitHubOutageStaysQuietAndAFaultIsCapturedOnce(t *testing.T) {
	c := mock_client.New()
	c.Failure = errors.New("github answered nonsense")
	r := memory.New()
	w, okay := daemon_tester.New(t, c, metric.New(), r).Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	assert.Count(t, 1, r.Events())
	c.Failure = refusedDial()

	for range 2 {
		w.Poll()
	}

	assert.Count(t, 1, r.Events())
}

func TestAMissingOwnerIsCaptured(t *testing.T) {
	c := mock_client.New()
	c.Failure = not_found.Format("Not Found")
	r := memory.New()
	w, okay := daemon_tester.New(t, c, metric.New(), r).Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	assert.Count(t, 1, r.Events())
}

func TestTheDaemonHandsOutTheGitHubGuardFirst(t *testing.T) {
	_, okay := daemon_tester.New(
		t,
		mock_client.New(),
		metric.New(),
		memory.New(),
	).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

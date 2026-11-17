package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestGitHubOutageStaysQuietAndFaultIsCapturedOnce(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = errors.New("github answered nonsense")
	w, okay := s.Daemon().Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	assert.Count(t, 1, s.Reporter.Events())
	s.Source.Failure = outage_tester.RefusedDial()

	for range 2 {
		w.Poll()
	}

	assert.Count(t, 1, s.Reporter.Events())
}

func TestMissingOwnerIsCaptured(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = not_found.Format("Not Found")
	w, okay := s.Daemon().Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	assert.Count(t, 1, s.Reporter.Events())
}

func TestDaemonHandsOutGitHubGuardFirst(t *testing.T) {
	_, okay := base.New(t).Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

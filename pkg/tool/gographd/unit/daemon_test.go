package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/teams/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gographd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gographd/option"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New(t)
	o := option.New()
	o.Source = func() face.TeamsClient { return c }
	d := daemon.New(o, logger.New(t.Context()), memory.New())
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

func TestCloseClosesTeamsClient(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Source.Closed)
}

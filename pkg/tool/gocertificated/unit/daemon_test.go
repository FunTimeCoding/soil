package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesForgeItBuilt(t *testing.T) {
	f := service_tester.NewForge(t)
	daemon.New(
		store_tester.NewOption(t, f),
		logger.New(t.Context()),
		memory.New(),
	).Close()
	assert.Integer(t, 1, f.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		store_tester.NewOption(t, mock_client.New()),
		logger.New(t.Context()),
		memory.New(),
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

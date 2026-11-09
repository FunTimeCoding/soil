package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/opnsense/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesClientItBuilt(t *testing.T) {
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.OpnsenseClient { return c }
	daemon.New(o, logger.New(t.Context()), memory.New()).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t, c).Close()
	goleak.VerifyNone(t, before)
}

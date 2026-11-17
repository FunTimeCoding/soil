package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesClientItBuilt(t *testing.T) {
	c := mock_client.New(t)
	o := option.New()
	o.Source = func() face.ItermClient { return c }
	daemon.New(o, logger.New(t.Context()), memory.New()).Close()
	assert.Integer(t, 1, c.Closed())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	before := goleak.IgnoreCurrent()
	s.Daemon().Close()
	goleak.VerifyNone(t, before)
}

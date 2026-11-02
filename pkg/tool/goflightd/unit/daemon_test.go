package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonSamplesWirelessWhenItAnswers(t *testing.T) {
	s := mock_system.New()
	t.Cleanup(s.Close)
	s.Answer(constant.Sudo, constant.WirelessSample)
	d := daemon.New(base.NewOption(t, s), logger.New(t.Context()), memory.New())
	t.Cleanup(d.Close)
	assert.Count(t, 4, d.Workers())
}

func TestTheDaemonSkipsWirelessWhenItFails(t *testing.T) {
	s := mock_system.New()
	t.Cleanup(s.Close)
	s.Fail(constant.Sudo, errors.Format("exit", 1))
	d := daemon.New(base.NewOption(t, s), logger.New(t.Context()), memory.New())
	t.Cleanup(d.Close)
	assert.Count(t, 3, d.Workers())
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	s := mock_system.New()
	t.Cleanup(s.Close)
	s.Serve(mock_stream.New(nil, true, errors.Format("killed", 9)))
	o := base.NewOption(t, s)
	before := goleak.IgnoreCurrent()
	d := daemon.New(o, logger.New(context.Background()), memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}

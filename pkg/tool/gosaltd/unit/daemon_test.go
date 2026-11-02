package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	o := newOption(t)
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

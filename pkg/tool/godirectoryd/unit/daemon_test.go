package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestWaitAnswersOnceDirectoryAnswers(t *testing.T) {
	assert.FatalOnError(t, base.New(t).Daemon().Wait(t.Context()))
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := base.New(t).Daemon()
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

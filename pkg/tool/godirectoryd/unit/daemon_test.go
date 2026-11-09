package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestWaitAnswersOnceDirectoryAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, service_tester.NewSource(), memory.New()).
			Wait(t.Context()),
	)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon_tester.New(t, service_tester.NewSource(), memory.New())
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

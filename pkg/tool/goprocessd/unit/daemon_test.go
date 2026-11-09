package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/unit/daemon_tester"
	"testing"
	"time"
)

func TestDaemonHandsItsSupervisorToLifecycle(t *testing.T) {
	d := daemon_tester.New(t, "alfa: sleep 60\n")
	workers := d.Workers()
	assert.Count(t, 1, workers)
	_, okay := workers[0].(*supervisor.Supervisor)
	assert.True(t, okay)
}

func TestWaitEndsWhenEveryProcessHasExited(t *testing.T) {
	done := make(chan struct{})
	returned := make(chan struct{})
	go func() {
		goprocessd.WaitForStop(done)
		close(returned)
	}()
	close(done)

	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("the wait did not end when done closed")
	}
}

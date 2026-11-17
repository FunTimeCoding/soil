package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/base"
	"testing"
	"time"
)

func TestCloseWaitsForRunInFlightToRecordItself(t *testing.T) {
	s := base.New(t)
	s.Hold()
	s.Start()
	s.WaitForRunning()
	closed := make(chan struct{})
	go func() {
		s.Daemon().Close()
		close(closed)
	}()

	select {
	case <-closed:
		t.Fatal("close returned with a run in flight")
	case <-time.After(200 * time.Millisecond):
	}

	s.Release()

	select {
	case <-closed:
	case <-time.After(10 * time.Second):
		t.Fatal("close did not return after the run finished")
	}

	runs := s.StoredRuns()
	assert.String(t, "success", runs[0].Status)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestStartAbandonsRunPreviousProcessLeftRunning(t *testing.T) {
	s := base.New(t)
	s.Hold()
	s.Start()
	s.WaitForRunning()
	s.Reload()
	runs := s.StoredRuns()
	assert.String(t, "error", runs[0].Status)
	assert.String(
		t,
		"abandoned: the daemon stopped before the run finished",
		runs[0].ErrorOutput,
	)
}

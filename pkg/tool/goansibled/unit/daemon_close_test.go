package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/daemon_tester"
	"testing"
	"time"
)

func TestCloseWaitsForRunInFlightToRecordItself(t *testing.T) {
	s, d := daemon_tester.New(t)
	s.Hold()
	daemon_tester.Start(d)
	s.WaitForRunning()
	closed := make(chan struct{})
	go func() {
		d.Close()
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
	s, d := daemon_tester.New(t)
	s.Hold()
	daemon_tester.Start(d)
	s.WaitForRunning()
	daemon_tester.Build(t, s)
	runs := s.StoredRuns()
	assert.String(t, "error", runs[0].Status)
	assert.String(
		t,
		"abandoned: the daemon stopped before the run finished",
		runs[0].ErrorOutput,
	)
}

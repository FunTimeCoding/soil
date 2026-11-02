package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"testing"
	"time"
)

func TestCloseWaitsForTheRunInFlightToRecordItself(t *testing.T) {
	s := ansible_tester.New(t, nil)
	s.WaitForRun()
	s.Hold()
	assert.FatalOnError(t, s.Service.Trigger(trigger.Request{}))
	s.WaitForRunning()
	closed := make(chan struct{})
	go func() {
		s.Service.Close()
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

func TestStartAbandonsARunAPreviousProcessLeftRunning(t *testing.T) {
	s := ansible_tester.New(t, nil)
	s.WaitForRun()
	s.Hold()
	assert.FatalOnError(t, s.Service.Trigger(trigger.Request{}))
	s.WaitForRunning()
	reopened := s.Reopen()
	runs, e := reopened.Recent(10)
	assert.FatalOnError(t, e)
	assert.String(t, "error", runs[0].Status)
	assert.String(
		t,
		"abandoned: the daemon stopped before the run finished",
		runs[0].ErrorOutput,
	)
}

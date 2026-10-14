package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleConstant "github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/unit/runner_tester"
	"testing"
	"time"
)

func TestStopWaitsForApplyInFlight(t *testing.T) {
	s := runner_tester.NewHeld(t)
	s.WaitForApply(1)
	s.Stop()
	s.ExpectRunning()
	s.Release()
	s.WaitForFinish()
	assert.Integer(t, 1, s.ApplyCount())
}

func TestStopAnswersTriggerQueuedDuringApply(t *testing.T) {
	for range 10 {
		s := runner_tester.NewHeld(t)
		s.WaitForApply(1)
		response := make(chan *trigger.Result, 1)
		s.Trigger(trigger.Request{Response: response})
		s.Stop()
		s.Release()
		s.WaitForFinish()
		assert.Integer(t, 1, s.ApplyCount())
		assert.Error(t, (<-response).Error)
	}
}

func TestInterruptStopsApplyProcess(t *testing.T) {
	s := runner_tester.NewSleeper(t)
	s.WaitForApply(1)
	s.Stop()
	s.ExpectRunning()
	s.Runner.Interrupt()
	s.WaitForFinish()
}

func TestStopTwice(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.Stop()
	s.Stop()
	s.WaitForFinish()
}

func TestLifecycleWaitsForApplyWithinGrace(t *testing.T) {
	s := runner_tester.NewHeld(t)
	s.WaitForApply(1)
	l := lifecycle.New(logger.New(context.Background())).WithWorker(s.Runner)
	start := time.Now()
	time.AfterFunc(200*time.Millisecond, s.Release)
	l.Close()
	s.ExpectFinished()
	assert.Greater(t, 200*time.Millisecond, time.Since(start))
	assert.Integer(t, 1, s.ApplyCount())
}

func TestLifecycleInterruptsApplyAtGrace(t *testing.T) {
	s := runner_tester.NewSleeper(t)
	s.WaitForApply(1)
	l := lifecycle.New(logger.New(context.Background())).WithGrace(
		lifecycleConstant.InterruptMargin + 200*time.Millisecond,
	).WithWorker(s.Runner)
	start := time.Now()
	l.Close()
	s.ExpectFinished()
	assert.Greater(t, 200*time.Millisecond, time.Since(start))
	assert.Less(t, 2*time.Second, time.Since(start))
}

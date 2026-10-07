package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/integration/runner_tester"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSetupReturnsFalseSkipsApply(t *testing.T) {
	applied := false
	r := runner.New(
		runner_option.Option{
			SetupFunction: func() bool { return false },
			ApplyFunction: func(
				_ map[string]any,
				_ string,
			) any {
				applied = true

				return nil
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	r.Start()
	defer r.Stop()
	time.Sleep(100 * time.Millisecond)
	assert.False(t, applied)
}

func TestDrainChannelsOnStop(t *testing.T) {
	gate := make(chan struct{})
	t.Cleanup(
		func() {
			select {
			case <-gate:
			default:
				close(gate)
			}
		},
	)
	r := runner.New(
		runner_option.Option{
			SetupFunction: func() bool {
				<-gate

				return false
			},
			ApplyFunction: func(
				_ map[string]any,
				_ string,
			) any {
				return nil
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	r.Start()
	response := make(chan *trigger.Result, 1)
	e := r.Trigger(trigger.Request{Response: response})
	assert.FatalOnError(t, e)
	close(gate)

	select {
	case result := <-response:
		assert.Error(t, result.Error)
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for drain response")
	}
}

func TestTriggerCallsApply(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.Trigger(trigger.Request{})
	s.WaitForApply(2)
	assert.String(t, "manual", s.LastApply().TriggerSource)
}

func TestTriggerWithParameters(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.Trigger(trigger.Request{Parameters: map[string]any{"target": "specific"}})
	s.WaitForApply(2)
	assert.String(t, "specific", s.LastApply().Parameters["target"].(string))
}

func TestSynchronousTrigger(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	response := make(chan *trigger.Result, 1)
	s.Trigger(trigger.Request{Response: response})
	result := <-response
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Value)
}

func TestCloneFilemodeOff(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	c := run.New()
	c.Directory = s.ClonePath
	c.Start("git", "config", "core.filemode")
	assert.String(t, "false", strings.TrimSpace(c.OutputString))
}

func TestSyncNoChanges(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	result := s.Sync()
	assert.False(t, result.Changed)
	assert.Nil(t, result.Error)
}

func TestSyncRemovesStaleIndexLock(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	lock := filepath.Join(s.ClonePath, constant.RunnerIndexLock)
	system.WriteFile(lock, nil, 0o644)
	s.PushCommit("tracked.txt", "after lock")
	result := s.Sync()
	assert.True(t, result.Changed)
	assert.Nil(t, result.Error)
	assert.False(t, system.FileExists(lock))
	assert.String(t, "after lock", system.ReadFile(s.ClonePath, "tracked.txt"))
}

func TestSyncHealsLocalDrift(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.PushCommit("tracked.txt", "first")
	result := s.Sync()
	assert.True(t, result.Changed)
	assert.Nil(t, result.Error)
	system.WriteFile(
		filepath.Join(s.ClonePath, "tracked.txt"),
		[]byte("drift"),
		0o755,
	)
	s.PushCommit("tracked.txt", "second")
	result = s.Sync()
	assert.True(t, result.Changed)
	assert.Nil(t, result.Error)
	assert.String(t, "second", system.ReadFile(s.ClonePath, "tracked.txt"))
}

func TestSyncHealsCorruptRepository(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	assert.FatalOnError(t, os.RemoveAll(filepath.Join(s.ClonePath, ".git")))
	first := s.Sync()
	assert.Error(t, first.Error)
	quarantined, e := filepath.Glob(join.Empty(s.ClonePath, ".quarantine.*"))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, quarantined)
	second := s.Sync()
	assert.Nil(t, second.Error)
}

func TestRepeatedFetchFailureForcesCloneOnValidRepository(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.BreakRemote()

	for range constant.RunnerHealThreshold {
		assert.Error(t, s.Sync().Error)
	}

	quarantined, e := filepath.Glob(join.Empty(s.ClonePath, ".quarantine.*"))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, quarantined)
	assert.Nil(t, s.Sync().Error)
}

func TestFetchFailureBelowThresholdKeepsRepository(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.BreakRemote()
	assert.Error(t, s.Sync().Error)
	quarantined, e := filepath.Glob(join.Empty(s.ClonePath, ".quarantine.*"))
	assert.FatalOnError(t, e)
	assert.Count(t, 0, quarantined)
}

package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/integration/runner_tester"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"testing"
)

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

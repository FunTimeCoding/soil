package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/integration_test/runner_tester"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestSyncNoChanges(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	result := s.Sync()
	assert.False(t, result.Changed)
	assert.Nil(t, result.Error)
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

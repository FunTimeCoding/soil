package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/runner/integration_test/runner_tester"
	"testing"
)

func TestSyncNoChanges(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	result := s.Sync()
	assert.False(t, result.Changed)
	assert.Nil(t, result.Error)
}

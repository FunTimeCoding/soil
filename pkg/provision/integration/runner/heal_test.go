package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/integration/runner_tester"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"testing"
)

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

package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/integration_test/runner_tester"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
	"testing"
)

func TestCloneFilemodeOff(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	c := run.New()
	c.Directory = s.ClonePath
	c.Start("git", "config", "core.filemode")
	assert.String(t, "false", strings.TrimSpace(c.OutputString))
}

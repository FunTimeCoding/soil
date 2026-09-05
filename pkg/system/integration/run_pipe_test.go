package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestRunPipe(t *testing.T) {
	windowsSkip(t)
	stdout, stderr := run.New().Pipe("pipe input\n", "cat")
	assert.String(t, "pipe input\n", stdout)
	assert.String(t, "", stderr)
}

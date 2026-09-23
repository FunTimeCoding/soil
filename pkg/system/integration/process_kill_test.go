package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestProcessKillIsNotAnError(t *testing.T) {
	windowsSkip(t)
	p := run.New().Open("sleep", "30")
	assert.Nil(t, p.Kill())
	assert.Integer(t, -1, p.ExitCode())
}

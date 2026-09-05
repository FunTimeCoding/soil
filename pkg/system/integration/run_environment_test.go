package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestRunEnvironment(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.Environment("INTEGRATION_VALUE", "expected")
	output := r.Start("sh", "-c", "echo $INTEGRATION_VALUE")
	assert.String(t, "expected\n", output)
}

func TestRunSetEnvironment(t *testing.T) {
	windowsSkip(t)
	r := run.New()
	r.SetEnvironment([]string{"INTEGRATION_ONLY=replaced"})
	output := r.Start("sh", "-c", "echo $INTEGRATION_ONLY:$HOME")
	assert.String(t, "replaced:\n", output)
}

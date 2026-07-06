package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os/exec"
	"testing"
)

func TestExitErrorCode(t *testing.T) {
	assert.Integer(
		t,
		-1,
		ExitErrorCode(error(&exec.ExitError{})),
	)
}

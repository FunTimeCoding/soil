package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"os/exec"
	"testing"
)

func TestExitErrorCode(t *testing.T) {
	assert.Integer(
		t,
		-1,
		system.ExitErrorCode(error(&exec.ExitError{})),
	)
}

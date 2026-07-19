package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/systemd/command"
	"testing"
)

func TestCommand(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"systemctl",
			"list-units",
			"--output",
			"json",
			"--state",
			"failed",
		},
		command.Failed(),
	)
}

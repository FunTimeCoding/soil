package ssh

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(
		t,
		"",
		environmentPrefix(command.New(upper.Alfa)),
	)
}

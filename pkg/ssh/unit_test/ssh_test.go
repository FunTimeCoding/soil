package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(
		t,
		"",
		ssh.EnvironmentPrefix(command.New(upper.Alfa)),
	)
}

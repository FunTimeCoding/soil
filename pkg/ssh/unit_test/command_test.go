package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestCommand(t *testing.T) {
	assert.NotNil(t, command.New(upper.Alfa))
}

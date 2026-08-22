package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func TestStartFailureCarriesCommandError(t *testing.T) {
	c := run.New().NoPanic()
	c.Start("sh", "-c", "echo oops >&2; exit 3")
	assert.NotNil(t, c.Error)
	assert.Integer(t, 3, c.Exit)
	assert.True(t, command.Is(c.Error))
	failure := c.Error.(*command.CommandError)
	assert.String(
		t,
		"sh -c echo oops >&2; exit 3: exit status 3",
		failure.Error(),
	)
	assert.StringContains(t, "oops", failure.Stderr)
}

//go:build local

package confluence

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestUserAndSpacesRead(t *testing.T) {
	c := confluence.NewEnvironment()
	u, e := c.User()
	assert.FatalOnError(t, e)
	assert.True(t, u != nil)
	spaces, f := c.Spaces()
	assert.FatalOnError(t, f)
	assert.NotEmpty(t, spaces)
}

func TestMissingPageIsNotFound(t *testing.T) {
	_, e := confluence.NewEnvironment().Page("999999999999")
	assert.True(t, not_found.Is(e))
}

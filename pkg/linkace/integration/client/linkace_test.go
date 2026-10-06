//go:build local

package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/linkace"
	"testing"
)

func TestListsAndTagsRead(t *testing.T) {
	c := linkace.NewEnvironment()
	lists, e := c.Lists()
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, lists)
	tags, f := c.Tags()
	assert.FatalOnError(t, f)
	assert.NotEmpty(t, tags)
}

func TestMissingLinkIsNotFound(t *testing.T) {
	_, e := linkace.NewEnvironment().LinkByIdentifier(999999999)
	assert.True(t, not_found.Is(e))
}

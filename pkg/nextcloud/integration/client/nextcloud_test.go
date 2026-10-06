//go:build local

package client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/nextcloud"
	"testing"
)

func TestStatusAndRootListing(t *testing.T) {
	c := nextcloud.NewEnvironment()
	assert.FatalOnError(t, c.Status())
	files, e := c.ReadDirectory("/")
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, files)
}

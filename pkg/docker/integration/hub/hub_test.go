//go:build local

package hub

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"testing"
)

func TestGolangTagsRead(t *testing.T) {
	owned := hub.New()
	defer owned.Close()
	tags, e := owned.Tags("library/golang")
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, tags)
}

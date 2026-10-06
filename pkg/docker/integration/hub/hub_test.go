//go:build local

package hub

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"testing"
)

func TestGolangTagsRead(t *testing.T) {
	tags, e := hub.New().Tags("library/golang")
	assert.FatalOnError(t, e)
	assert.NotEmpty(t, tags)
}

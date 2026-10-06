package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestInsideDirectory(t *testing.T) {
	assert.True(t, system.InsideDirectory("/src/user", "/src/user/pkg/a.go"))
	assert.True(t, system.InsideDirectory("/src/user/", "/src/user/a.go"))
	assert.False(t, system.InsideDirectory("/src/user", "/src/library/a.go"))
	assert.False(t, system.InsideDirectory("/src/user", "/src/user-other/a.go"))
	assert.False(t, system.InsideDirectory("/src/user", "relative/a.go"))
}

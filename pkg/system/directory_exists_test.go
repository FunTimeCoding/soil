package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDirectoryExists(t *testing.T) {
	assert.True(t, DirectoryExists(WorkDirectory()))
	assert.False(t, DirectoryExists("non-existent-directory"))
}

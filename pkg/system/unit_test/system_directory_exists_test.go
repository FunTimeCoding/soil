package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestSystemDirectoryExists(t *testing.T) {
	assert.True(t, system.DirectoryExists(system.WorkDirectory()))
	assert.False(t, system.DirectoryExists("non-existent-directory"))
}

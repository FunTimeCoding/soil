//go:build !windows

package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	assert.NotNil(t, DiskUsage(WorkDirectory()))
}

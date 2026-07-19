//go:build !windows

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	assert.NotNil(t, system.DiskUsage(system.WorkDirectory()))
}

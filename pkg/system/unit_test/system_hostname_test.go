package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestHostname(t *testing.T) {
	assert.True(t, len(system.Hostname()) > 0)
}

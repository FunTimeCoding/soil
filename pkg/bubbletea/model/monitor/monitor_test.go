package monitor

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMonitor(t *testing.T) {
	assert.NotNil(t, New(false))
}

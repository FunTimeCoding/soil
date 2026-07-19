package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"testing"
)

func TestMonitor(t *testing.T) {
	assert.NotNil(t, monitor.New(false))
}

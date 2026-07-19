package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	"testing"
)

func TestTick(t *testing.T) {
	assert.NotNil(t, tick.Command())
}

package tick

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestTick(t *testing.T) {
	assert.NotNil(t, Command())
}

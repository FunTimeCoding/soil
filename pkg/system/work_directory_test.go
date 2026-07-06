package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"strings"
	"testing"
)

func TestWorkDirectory(t *testing.T) {
	assert.True(
		t,
		strings.Contains(WorkDirectory(), "system"),
	)
}

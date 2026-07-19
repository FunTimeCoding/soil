package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
	"testing"
)

func TestWorkDirectory(t *testing.T) {
	assert.True(
		t,
		strings.Contains(system.WorkDirectory(), "system"),
	)
}

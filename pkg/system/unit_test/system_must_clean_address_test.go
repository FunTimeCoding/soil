package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestCleanAddress(t *testing.T) {
	assert.String(t, "0.0.0.0", system.MustCleanAddress("0.0.0.0:22"))
	assert.String(t, "0.0.0.0", system.MustCleanAddress("0.0.0.0/32"))
}

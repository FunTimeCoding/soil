package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCleanAddress(t *testing.T) {
	assert.String(t, "0.0.0.0", MustCleanAddress("0.0.0.0:22"))
	assert.String(t, "0.0.0.0", MustCleanAddress("0.0.0.0/32"))
}

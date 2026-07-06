package netbox

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestWithVerbose(t *testing.T) {
	assert.NotNil(t, WithVerbose(true))
}

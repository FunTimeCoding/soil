package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"testing"
)

func TestWithVerbose(t *testing.T) {
	assert.NotNil(t, netbox.WithVerbose(true))
}

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRearPort(t *testing.T) {
	assert.NotNil(t, rear_port.New(&netbox.RearPort{}))
}

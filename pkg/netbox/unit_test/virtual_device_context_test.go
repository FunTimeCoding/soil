package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestVirtualDeviceContext(t *testing.T) {
	assert.NotNil(
		t,
		virtual_device_context.New(&netbox.VirtualDeviceContext{}),
	)
}

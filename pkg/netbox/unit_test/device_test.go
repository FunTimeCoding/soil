package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDevice(t *testing.T) {
	assert.NotNil(t, device.New(&netbox.DeviceWithConfigContext{}))
}

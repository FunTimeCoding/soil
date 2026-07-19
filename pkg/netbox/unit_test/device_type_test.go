package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDeviceType(t *testing.T) {
	assert.NotNil(t, device_type.New(&netbox.DeviceType{}))
}

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDeviceRole(t *testing.T) {
	assert.NotNil(t, device_role.New(&netbox.DeviceRole{}))
}

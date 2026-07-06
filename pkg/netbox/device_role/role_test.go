package device_role

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRole(t *testing.T) {
	assert.NotNil(t, New(&netbox.DeviceRole{}))
}

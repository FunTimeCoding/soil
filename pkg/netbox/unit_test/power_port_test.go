package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestPowerPort(t *testing.T) {
	assert.NotNil(t, power_port.New(&netbox.PowerPort{}))
}

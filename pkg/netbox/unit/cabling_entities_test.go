package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestCablingEntities(t *testing.T) {
	assert.NotNil(t, cable.New(&netbox.Cable{}))
	assert.NotNil(t, console_port.New(&netbox.ConsolePort{}))
	assert.NotNil(t, console_server_port.New(&netbox.ConsoleServerPort{}))
	assert.NotNil(t, front_port.New(&netbox.FrontPort{}))
	assert.NotNil(t, network.New(&netbox.Interface{}))
	assert.NotNil(
		t,
		physical_address.New(&netbox.MACAddress{Display: soil.PhysicalTest0}),
	)
	assert.NotNil(t, rear_port.New(&netbox.RearPort{}))
}

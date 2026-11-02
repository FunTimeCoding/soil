package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/module"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestDeviceEntities(t *testing.T) {
	assert.NotNil(t, device_bay_template.New(&netbox.DeviceBayTemplate{}))
	assert.NotNil(t, device_role.New(&netbox.DeviceRole{}))
	assert.NotNil(t, device.New(&netbox.DeviceWithConfigContext{}))
	assert.NotNil(t, device_type.New(&netbox.DeviceType{}))
	assert.NotNil(t, inventory_item_role.New(&netbox.InventoryItemRole{}))
	assert.NotNil(t, inventory_item.New(&netbox.InventoryItem{}))
	assert.NotNil(t, manufacturer.New(&netbox.Manufacturer{}))
	assert.NotNil(t, module_bay.New(&netbox.ModuleBay{}))
	assert.NotNil(t, module.New(&netbox.Module{}))
	assert.NotNil(t, module_type_profile.New(&netbox.ModuleTypeProfile{}))
	assert.NotNil(t, module_type.New(&netbox.ModuleType{}))
	assert.NotNil(t, platform.New(&netbox.Platform{}))
	assert.NotNil(t, virtual_chassis.New(&netbox.VirtualChassis{}))
	assert.NotNil(t, virtual_device_context.New(&netbox.VirtualDeviceContext{}))
}

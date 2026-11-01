package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
	"github.com/funtimecoding/soil/pkg/netbox/module"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
)

type ModuleClient interface {
	DeviceBayTemplates() ([]*device_bay_template.Template, error)
	DeviceModuleBays(device string) ([]*module_bay.Bay, error)
	InventoryItemRoles() ([]*inventory_item_role.Role, error)
	InventoryItems() ([]*inventory_item.Item, error)
	ModuleBays() ([]*module_bay.Bay, error)
	ModuleTypeProfiles() ([]*module_type_profile.Profile, error)
	ModuleTypes() ([]*module_type.Type, error)
	Modules() ([]*module.Module, error)
	MustDeviceBayTemplates() []*device_bay_template.Template
	MustDeviceModuleBays(device string) []*module_bay.Bay
	MustInventoryItemRoles() []*inventory_item_role.Role
	MustInventoryItems() []*inventory_item.Item
	MustModuleBays() []*module_bay.Bay
	MustModuleTypeProfiles() []*module_type_profile.Profile
	MustModuleTypes() []*module_type.Type
	MustModules() []*module.Module
	MustVirtualChassis() []*virtual_chassis.Chassis
	MustVirtualDeviceContexts() []*virtual_device_context.Context
	VirtualChassis() ([]*virtual_chassis.Chassis, error)
	VirtualDeviceContexts() ([]*virtual_device_context.Context, error)
}

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_choice"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/netbox/module"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
	"github.com/funtimecoding/soil/pkg/netbox/module_type_profile"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
	"github.com/funtimecoding/soil/pkg/netbox/service"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/source"
	"github.com/funtimecoding/soil/pkg/netbox/system_number"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/tenant_group"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/netbox/user"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestEntities(t *testing.T) {
	assert.NotNil(t, cable.New(&netbox.Cable{}))
	assert.NotNil(t, cluster_group.New(&netbox.ClusterGroup{}))
	assert.NotNil(t, cluster.New(&netbox.Cluster{}))
	assert.NotNil(t, cluster_type.New(&netbox.ClusterType{}))
	assert.NotNil(t, configuration_context.New(&netbox.ConfigContext{}))
	assert.NotNil(t, configuration_template.New(&netbox.ConfigTemplate{}))
	assert.NotNil(t, console_port.New(&netbox.ConsolePort{}))
	assert.NotNil(t, console_server_port.New(&netbox.ConsoleServerPort{}))
	assert.NotNil(t, contact_group.New(&netbox.ContactGroup{}))
	assert.NotNil(t, contact_role.New(&netbox.ContactRole{}))
	assert.NotNil(t, contact.New(&netbox.Contact{}))
	assert.NotNil(t, custom_field_choice.New(&netbox.CustomFieldChoiceSet{}))
	assert.NotNil(t, custom_field.New(&netbox.CustomField{}))
	assert.NotNil(t, custom_link.New(&netbox.CustomLink{}))
	assert.NotNil(t, device_bay_template.New(&netbox.DeviceBayTemplate{}))
	assert.NotNil(t, device_role.New(&netbox.DeviceRole{}))
	assert.NotNil(t, device.New(&netbox.DeviceWithConfigContext{}))
	assert.NotNil(t, device_type.New(&netbox.DeviceType{}))
	assert.NotNil(t, export_template.New(&netbox.ExportTemplate{}))
	assert.NotNil(t, front_port.New(&netbox.FrontPort{}))
	assert.NotNil(
		t,
		internet_address.New(&netbox.IPAddress{Display: "10.0.0.1/24"}),
	)
	assert.NotNil(t, inventory_item_role.New(&netbox.InventoryItemRole{}))
	assert.NotNil(t, inventory_item.New(&netbox.InventoryItem{}))
	assert.NotNil(t, location.New(&netbox.Location{}))
	assert.NotNil(t, manufacturer.New(&netbox.Manufacturer{}))
	assert.NotNil(t, module_bay.New(&netbox.ModuleBay{}))
	assert.NotNil(t, module.New(&netbox.Module{}))
	assert.NotNil(t, module_type_profile.New(&netbox.ModuleTypeProfile{}))
	assert.NotNil(t, module_type.New(&netbox.ModuleType{}))
	assert.NotNil(t, network.New(&netbox.Interface{}))
	assert.NotNil(t, notification_group.New(&netbox.NotificationGroup{}))
	assert.NotNil(
		t,
		physical_address.New(&netbox.MACAddress{Display: constant.PhysicalTest0}),
	)
	assert.NotNil(t, platform.New(&netbox.Platform{}))
	assert.NotNil(t, power_feed.New(&netbox.PowerFeed{}))
	assert.NotNil(t, power_outlet.New(&netbox.PowerOutlet{}))
	assert.NotNil(t, power_panel.New(&netbox.PowerPanel{}))
	assert.NotNil(t, power_port.New(&netbox.PowerPort{}))
	assert.NotNil(t, prefix.New(&netbox.Prefix{}))
	assert.NotNil(t, rack_role.New(&netbox.RackRole{}))
	assert.NotNil(t, rack.New(&netbox.Rack{}))
	assert.NotNil(t, rack_type.New(&netbox.RackType{}))
	assert.NotNil(t, rear_port.New(&netbox.RearPort{}))
	assert.NotNil(t, reservation.New(&netbox.RackReservation{}))
	assert.NotNil(t, service_template.New(&netbox.ServiceTemplate{}))
	assert.NotNil(t, service.New(&netbox.Service{}))
	assert.NotNil(t, site.New(&netbox.Site{}))
	assert.NotNil(t, source.New(&netbox.DataSource{}))
	assert.NotNil(t, system_number.New(&netbox.ASN{}))
	assert.NotNil(t, tag.New(&netbox.Tag{}))
	assert.NotNil(t, tenant_group.New(&netbox.TenantGroup{}))
	assert.NotNil(t, tenant.New(&netbox.Tenant{}))
	assert.NotNil(t, tunnel_group.New(&netbox.TunnelGroup{}))
	assert.NotNil(t, tunnel.New(&netbox.Tunnel{}))
	assert.NotNil(t, user_group.New(&netbox.Group{}))
	assert.NotNil(t, user.New(&netbox.User{}))
	assert.NotNil(t, virtual_chassis.New(&netbox.VirtualChassis{}))
	assert.NotNil(
		t,
		virtual_device_context.New(&netbox.VirtualDeviceContext{}),
	)
	assert.NotNil(
		t,
		virtual_machine.New(&netbox.VirtualMachineWithConfigContext{}),
	)
	assert.NotNil(t, virtual_network_group.New(&netbox.VLANGroup{}))
	assert.NotNil(t, virtual_network.New(&netbox.VLAN{}))
	assert.NotNil(t, wireless_link.New(&netbox.WirelessLink{}))
	assert.NotNil(t, wireless_network_group.New(&netbox.WirelessLANGroup{}))
	assert.NotNil(t, wireless_network.New(&netbox.WirelessLAN{}))
}

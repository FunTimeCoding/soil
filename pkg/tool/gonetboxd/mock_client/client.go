package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

type Client struct {
	identifier int32
	devices    []*device.Device
	machines   []*virtual_machine.Machine
	roles      []*device_role.Role
	types      []*device_type.Type
	sites      []*site.Site
	tenants    []*tenant.Tenant
	interfaces map[int32][]*network.Interface
	contexts   []*configuration_context.Context
	sets       []*choice_set.Set
	fields     []*custom_field.Field
	calls      []string
}

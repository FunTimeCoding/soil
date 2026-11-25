package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestInternetAddressReadsDeviceHolder(t *testing.T) {
	v := netbox.NewIPAddressWithDefaults()
	v.SetDisplay("192.0.2.5/32")
	v.SetAssignedObjectType(constant.InterfaceAddress)
	v.AssignedObject = map[string]any{
		"name":   "eth0",
		"device": map[string]any{"name": "alfa"},
	}
	role := netbox.IPADDRESSROLEVALUE_VIP
	v.SetRole(netbox.IPAddressRole{Value: &role})
	status := netbox.IPADDRESSSTATUSVALUE_ACTIVE
	v.SetStatus(netbox.IPAddressStatus{Value: &status})
	v.SetDnsName("bravo.example.org")
	v.SetDescription("charlie")
	a := internet_address.New(v)
	assert.String(t, "eth0", a.Interface)
	assert.String(t, "alfa", a.Device)
	assert.String(t, "", a.VirtualMachine)
	assert.String(t, "vip", a.Role)
	assert.String(t, "active", a.Status)
	assert.String(t, "bravo.example.org", a.DomainName)
	assert.String(t, "charlie", a.Description)
}

func TestInternetAddressReadsMachineHolder(t *testing.T) {
	v := netbox.NewIPAddressWithDefaults()
	v.SetDisplay("192.0.2.9/32")
	v.SetAssignedObjectType(constant.VirtualInterfaceAddress)
	v.AssignedObject = map[string]any{
		"name":            "ens3",
		"virtual_machine": map[string]any{"name": "delta"},
	}
	a := internet_address.New(v)
	assert.String(t, "ens3", a.Interface)
	assert.String(t, "delta", a.VirtualMachine)
	assert.String(t, "", a.Device)
	assert.String(t, "", a.Role)
}

func TestInternetAddressWithoutHolderStaysEmpty(t *testing.T) {
	v := netbox.NewIPAddressWithDefaults()
	v.SetDisplay("192.0.2.1/32")
	a := internet_address.New(v)
	assert.String(t, "", a.Interface)
	assert.String(t, "", a.Device)
	assert.String(t, "", a.VirtualMachine)
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/object_type"
	"testing"
)

func TestLabelNamesKnownObjectType(t *testing.T) {
	assert.String(t, "Device", object_type.Label(constant.DeviceAddress))
	assert.String(t, "Prefix", object_type.Label(constant.PrefixAddress))
}

func TestLabelNamesCompoundObjectType(t *testing.T) {
	assert.String(t, "IP Address", object_type.Label("ipam.ipaddress"))
	assert.String(t, "VLAN", object_type.Label("ipam.vlan"))
	assert.String(
		t,
		"Virtual Machine",
		object_type.Label(constant.VirtualMachineAddress),
	)
}

func TestLabelKeepsUnknownObjectTypeRaw(t *testing.T) {
	assert.String(t, "dcim.cable", object_type.Label("dcim.cable"))
	assert.String(t, "", object_type.Label(""))
}

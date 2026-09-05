package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"net"
	"testing"
)

func TestDerive(t *testing.T) {
	result, e := network_device.Derive(0, 106)
	assert.Nil(t, e)
	assert.String(t, "02:00:00:00:00:6a", result)
}

func TestDeriveSeparatesInstances(t *testing.T) {
	first, e := network_device.Derive(0, 100)
	assert.Nil(t, e)
	second, f := network_device.Derive(1, 100)
	assert.Nil(t, f)
	assert.True(t, first != second)
}

func TestDerivedAddressIsLocallyAdministered(t *testing.T) {
	result, e := network_device.Derive(3, 999999)
	assert.Nil(t, e)
	address, f := net.ParseMAC(result)
	assert.Nil(t, f)
	assert.Integer(t, 2, int(address[0]&0x02))
	assert.Integer(t, 0, int(address[0]&0x01))
}

func TestDeriveRejectsOutOfRange(t *testing.T) {
	_, e := network_device.Derive(256, 100)
	assert.Error(t, e)
	_, f := network_device.Derive(0, -1)
	assert.Error(t, f)
}

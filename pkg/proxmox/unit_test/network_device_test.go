package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"testing"
)

func TestMachineDevice(t *testing.T) {
	d := network_device.New("net0", "virtio=02:00:00:00:00:01,bridge=vmbr0")
	assert.String(t, "net0", d.Name)
	assert.String(t, "virtio", d.Model)
	assert.String(t, "02:00:00:00:00:01", d.HardwareAddress)
	assert.String(t, "vmbr0", d.Bridge)
	assert.Integer(t, 0, d.Vlan)
}

func TestContainerDevice(t *testing.T) {
	d := network_device.New(
		"net0",
		"name=eth0,bridge=vmbr0,firewall=1,hwaddr=02:00:00:00:00:02,ip=dhcp,type=veth,tag=15",
	)
	assert.String(t, "eth0", d.Interface)
	assert.String(t, "02:00:00:00:00:02", d.HardwareAddress)
	assert.String(t, "vmbr0", d.Bridge)
	assert.Integer(t, 15, d.Vlan)
	assert.String(t, "", d.Model)
}

func TestDeviceIgnoresUnknownOptions(t *testing.T) {
	d := network_device.New(
		"net1",
		"virtio=02:00:00:00:00:03,bridge=vmbr1,firewall=1,mtu=9000",
	)
	assert.String(t, "virtio", d.Model)
	assert.String(t, "02:00:00:00:00:03", d.HardwareAddress)
	assert.String(t, "vmbr1", d.Bridge)
}

func TestDeviceSliceSortsNumerically(t *testing.T) {
	result := network_device.NewSlice(
		map[string]string{
			"net10": "virtio=02:00:00:00:00:10,bridge=vmbr0",
			"net2":  "virtio=02:00:00:00:00:02,bridge=vmbr0",
			"net0":  "virtio=02:00:00:00:00:00,bridge=vmbr0",
		},
	)
	assert.Integer(t, 3, len(result))
	assert.String(t, "net0", result[0].Name)
	assert.String(t, "net2", result[1].Name)
	assert.String(t, "net10", result[2].Name)
}

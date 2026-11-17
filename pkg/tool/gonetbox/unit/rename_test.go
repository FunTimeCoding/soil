package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"testing"
)

func TestUpdateDeviceNameFlags(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-device",
		"alfa",
		"--name",
		"bravo",
		"--primary-address",
		"192.0.2.1/24",
	)
	assert.Integer(t, 0, code)
	assert.Any(t, new("bravo"), c.Updated["alfa"].Name)
	assert.Any(t, new("192.0.2.1/24"), c.Updated["alfa"].PrimaryAddress)
}

func TestUpdateVirtualMachineNameFlags(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-virtual-machine",
		"alfa",
		"--name",
		"bravo",
		"--primary-address",
		"192.0.2.1/24",
	)
	assert.Integer(t, 0, code)
	assert.Any(t, new("bravo"), c.UpdatedMachines["alfa"].Name)
	assert.Any(t, new("192.0.2.1/24"), c.UpdatedMachines["alfa"].PrimaryAddress)
}

func TestUpdateDeviceRoleNameFlag(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-device-role",
		"alfa",
		"--name",
		"bravo",
	)
	assert.Integer(t, 0, code)
	assert.Any(
		t,
		client.UpdateDeviceRoleRequest{Name: new("bravo")},
		c.UpdatedRoles["alfa"],
	)
}

func TestUpdateDeviceRoleWithoutNameSendsNone(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(t, c, "update-device-role", "alfa")
	assert.Integer(t, 0, code)
	assert.Any(t, client.UpdateDeviceRoleRequest{}, c.UpdatedRoles["alfa"])
}

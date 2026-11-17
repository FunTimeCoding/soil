package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func TestUpdateDeviceLocalContextFlagDecoded(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-device",
		"alfa",
		"--local-context",
		`{"bravo": {"charlie": "delta"}}`,
	)
	assert.Integer(t, 0, code)
	assert.Any(
		t,
		&map[string]any{"bravo": map[string]any{"charlie": "delta"}},
		c.Updated["alfa"].LocalContext,
	)
}

func TestUpdateDeviceWithoutLocalContextFlagSendsNone(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-device",
		"alfa",
		"--description",
		"bravo",
	)
	assert.Integer(t, 0, code)
	assert.Nil(t, c.Updated["alfa"].LocalContext)
}

func TestUpdateVirtualMachineEmptyLocalContextFlagSendsEmptyObject(
	t *testing.T,
) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-virtual-machine",
		"alfa",
		"--local-context",
		"{}",
	)
	assert.Integer(t, 0, code)
	assert.Any(t, &map[string]any{}, c.UpdatedMachines["alfa"].LocalContext)
}

func TestMalformedLocalContextFlagPanicsBeforeRequest(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	assert.PanicContains(
		t,
		"invalid character",
		func() {
			command_tester.Run(
				t,
				c,
				"update-device",
				"alfa",
				"--local-context",
				"bravo",
			)
		},
	)
	assert.Count(t, 0, c.Updated)
}

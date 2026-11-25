package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"testing"
)

func TestRemovalPrintsNothingOnSuccess(t *testing.T) {
	for _, arguments := range command_tester.Removals() {
		c := mock_client.New(t)
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() {
				output, failure, code = command_tester.New(t).Run(
					c,
					arguments...,
				)
			},
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
		c.Close()
	}
}

func TestFailedRemovalExitsWithItsMessage(t *testing.T) {
	for _, arguments := range command_tester.Removals() {
		c := mock_client.New(t)
		c.Failure = errors.New("foxtrot")
		_, failure, code := command_tester.New(t).Run(c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
		c.Close()
	}
}

func TestListPrefixAddressesEmitsHoldersInsidePrefix(t *testing.T) {
	c := mock_client.New(t)
	alfa := "alfa"
	bravo := "bravo"
	c.AddressList = []*client.Address{
		{Address: "192.0.2.5", Device: &alfa},
		{Address: "198.51.100.1", Device: &bravo},
	}
	output, _, code := command_tester.New(t).Run(
		c,
		"list-prefix-addresses",
		"192.0.2.0/24",
	)
	assert.Integer(t, 0, code)
	var a []client.Address
	notation.MustDecode(output, &a, false)
	assert.Count(t, 1, a)
	assert.String(t, "alfa", *a[0].Device)
}

func TestListDevicesEmitsDevices(t *testing.T) {
	c := mock_client.New(t)
	c.DeviceList = []*client.Device{{Identifier: 7, Name: "alfa"}}
	output, _, code := command_tester.New(t).Run(c, "list-devices")
	assert.Integer(t, 0, code)
	var d []client.Device
	notation.MustDecode(output, &d, false)
	assert.Count(t, 1, d)
	assert.String(t, "alfa", d[0].Name)
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"testing"
)

func TestSetInterfaceEthernetPowerCarriesType(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"set-interface-ethernet-power",
		"alfa",
		"bravo",
		"pse",
		"type2-ieee802.3at",
	)
	assert.Integer(t, 0, code)
	assert.Any(
		t,
		&client.EthernetPowerRequest{
			Mode: "pse",
			Type: new("type2-ieee802.3at"),
		},
		c.EthernetPower,
	)
}

func TestSetInterfaceEthernetPowerWithoutTypeSendsNone(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"set-interface-ethernet-power",
		"alfa",
		"bravo",
		"pd",
	)
	assert.Integer(t, 0, code)
	assert.Any(t, &client.EthernetPowerRequest{Mode: "pd"}, c.EthernetPower)
}

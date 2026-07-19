package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNetworkFindByName(t *testing.T) {
	i := network.New(
		&netbox.Interface{
			Name: network.Eth0,
			Type: netbox.InterfaceType{
				Value: new(
					netbox.InterfaceTypeValue(
						network.InterfaceVirtual,
					),
				),
			},
		},
	)
	interfaces := []*network.Interface{i}
	// Happy path
	assert.Any(
		t,
		i,
		network.FindByName(interfaces, network.Eth0),
	)
	// Not found
	var expected *network.Interface
	assert.Any(t, expected, network.FindByName(interfaces, network.Eth1))
}

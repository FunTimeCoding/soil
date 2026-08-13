package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNetworkFindByName(t *testing.T) {
	i := network.New(
		&netbox.Interface{
			Name: constant.Eth0,
			Type: netbox.InterfaceType{
				Value: new(netbox.InterfaceTypeValue(constant.InterfaceVirtual)),
			},
		},
	)
	interfaces := []*network.Interface{i}
	// Happy path
	assert.Any(t, i, network.FindByName(interfaces, constant.Eth0))
	// Not found
	var expected *network.Interface
	assert.Any(t, expected, network.FindByName(interfaces, constant.Eth1))
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"testing"
)

func TestNetworkFindByName(t *testing.T) {
	i := newVirtualInterface()
	assert.Any(t, i, network.FindByName([]*network.Interface{i}, constant.Eth0))
}

func TestNetworkFindByNameMissesUnknown(t *testing.T) {
	var expected *network.Interface
	assert.Any(
		t,
		expected,
		network.FindByName(
			[]*network.Interface{newVirtualInterface()},
			constant.Eth1,
		),
	)
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNetworkEntities(t *testing.T) {
	assert.NotNil(t, tunnel_group.New(&netbox.TunnelGroup{}))
	assert.NotNil(t, tunnel.New(&netbox.Tunnel{}))
	assert.NotNil(t, virtual_network_group.New(&netbox.VLANGroup{}))
	assert.NotNil(t, virtual_network.New(&netbox.VLAN{}))
	assert.NotNil(t, wireless_link.New(&netbox.WirelessLink{}))
	assert.NotNil(t, wireless_network_group.New(&netbox.WirelessLANGroup{}))
	assert.NotNil(t, wireless_network.New(&netbox.WirelessLAN{}))
}

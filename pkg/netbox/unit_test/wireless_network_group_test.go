package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestWirelessNetworkGroup(t *testing.T) {
	assert.NotNil(t, wireless_network_group.New(&netbox.WirelessLANGroup{}))
}

package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestVirtualNetworkGroup(t *testing.T) {
	assert.NotNil(t, virtual_network_group.New(&netbox.VLANGroup{}))
}

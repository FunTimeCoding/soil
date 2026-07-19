package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestVirtualNetwork(t *testing.T) {
	assert.NotNil(t, virtual_network.New(&netbox.VLAN{}))
}

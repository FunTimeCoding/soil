package wireless_network_group

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestGroup(t *testing.T) {
	assert.NotNil(t, New(&netbox.WirelessLANGroup{}))
}

package virtual_network

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestNetwork(t *testing.T) {
	assert.NotNil(t, New(&netbox.VLAN{}))
}

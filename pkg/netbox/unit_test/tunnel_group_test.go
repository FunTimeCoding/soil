package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTunnelGroup(t *testing.T) {
	assert.NotNil(t, tunnel_group.New(&netbox.TunnelGroup{}))
}

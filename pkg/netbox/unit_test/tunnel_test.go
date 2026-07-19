package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestTunnel(t *testing.T) {
	assert.NotNil(t, tunnel.New(&netbox.Tunnel{}))
}

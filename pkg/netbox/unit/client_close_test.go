package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net"
	"testing"
)

func TestClientCloseDropsIdleConnections(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureMacList),
		upstream_server.WithObject(),
	)
	c := netbox_tester.New(t, s.URL)
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.HardwarePhysicalAddresses(a)
	assert.FatalOnError(t, e)
	c.Close()
	_, e = c.HardwarePhysicalAddresses(a)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, s.Opened())
}

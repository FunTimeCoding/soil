package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/network"
	"testing"
)

func TestUpdateVirtualInterfaceSetsThePrimaryAddress(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	_, e := c.UpdateVirtualInterface(
		newAlfa(),
		"ens3",
		network.PhysicalAddress(constant.FixturePhysicalAddress),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Unmatched)
	assert.String(t, "set", r.Primary)
}

func TestUpdateVirtualInterfaceAssignsToTheVirtualInterface(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	_, e := c.UpdateVirtualInterface(
		newAlfa(),
		"ens3",
		network.PhysicalAddress(constant.FixturePhysicalAddress),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "virtualization.vminterface", r.Assigned)
}

package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestRackEntities(t *testing.T) {
	assert.NotNil(t, location.New(&netbox.Location{}))
	assert.NotNil(t, power_feed.New(&netbox.PowerFeed{}))
	assert.NotNil(t, power_outlet.New(&netbox.PowerOutlet{}))
	assert.NotNil(t, power_panel.New(&netbox.PowerPanel{}))
	assert.NotNil(t, power_port.New(&netbox.PowerPort{}))
	assert.NotNil(t, rack_role.New(&netbox.RackRole{}))
	assert.NotNil(t, rack.New(&netbox.Rack{}))
	assert.NotNil(t, rack_type.New(&netbox.RackType{}))
	assert.NotNil(t, reservation.New(&netbox.RackReservation{}))
	assert.NotNil(t, site.New(&netbox.Site{}))
}

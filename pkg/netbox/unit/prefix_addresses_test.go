package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestPrefixAddressesAskNetBoxForParent(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithObject(),
		upstream_server.WithAnswer(
			`{"count":1,"results":[{"id":5,"url":"https://netbox.example.org/api/ipam/ip-addresses/5/","display":"192.0.2.5/32","family":{"value":4,"label":"IPv4"},"address":"192.0.2.5/32","nat_outside":[],"assigned_object_type":"dcim.interface","assigned_object_id":7,"assigned_object":{"id":7,"name":"eth0","device":{"id":3,"name":"alfa"}}}]}`,
		),
	)
	v, e := netbox_tester.New(t, s.URL).PrefixAddresses("192.0.2.0/24")
	assert.FatalOnError(t, e)
	assert.String(t, "192.0.2.0/24", s.Last().URL.Query().Get("parent"))
	assert.Count(t, 1, v)
	assert.String(t, "alfa", v[0].Device)
	assert.String(t, "eth0", v[0].Interface)
}

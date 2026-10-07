package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/response"
	"github.com/funtimecoding/soil/pkg/technitium/zone"
)

func (c *Client) Zones() ([]*zone.Zone, error) {
	var result response.Zones

	return result.Zones, c.get("/zones/list", &result)
}

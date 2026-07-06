package hetzner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (c *Client) CreateRecord(
	z *zone.Zone,
	name string,
	recordType string,
	value string,
) {
	_, _, e := c.client.Zone.CreateRRSet(
		c.context,
		z.Raw,
		hcloud.ZoneRRSetCreateOpts{
			Name:    name,
			Type:    hcloud.ZoneRRSetType(recordType),
			Records: []hcloud.ZoneRRSetRecord{{Value: value}},
		},
	)
	errors.PanicOnError(e)
}

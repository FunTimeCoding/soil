package hetzner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (c *Client) DeleteRecord(
	z *zone.Zone,
	name string,
	recordType string,
) {
	rrset, _, e := c.client.Zone.GetRRSetByNameAndType(
		c.context,
		z.Raw,
		name,
		hcloud.ZoneRRSetType(recordType),
	)
	errors.PanicOnError(e)
	_, _, f := c.client.Zone.DeleteRRSet(c.context, rrset)
	errors.PanicOnError(f)
}

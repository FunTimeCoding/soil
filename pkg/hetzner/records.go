package hetzner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/hetzner/record"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (c *Client) Records(z *zone.Zone) []*record.Record {
	result, e := c.client.Zone.AllRRSetsWithOpts(
		c.context,
		z.Raw,
		hcloud.ZoneRRSetListOpts{},
	)
	errors.PanicOnError(e)

	return record.NewSlice(result, z)
}

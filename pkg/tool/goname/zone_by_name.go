package goname

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"os"
)

func zoneByName(
	c *hetzner.Client,
	name string,
) *zone.Zone {
	for _, z := range c.Zones() {
		if z.Name == name {
			return z
		}
	}

	console.Format("zone not found: %s\n", name)
	os.Exit(1)

	return nil
}

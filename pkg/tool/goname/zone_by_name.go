package goname

import (
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/funtimecoding/soil/pkg/terminal"
)

func zoneByName(
	c *hetzner.Client,
	name string,
	t *terminal.Terminal,
) *zone.Zone {
	for _, z := range c.Zones() {
		if z.Name == name {
			return z
		}
	}

	t.Exitf("zone not found: %s\n", name)

	return nil
}

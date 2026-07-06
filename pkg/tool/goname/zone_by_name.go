package goname

import (
	"fmt"
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

	fmt.Printf("zone not found: %s\n", name)
	os.Exit(1)

	return nil
}

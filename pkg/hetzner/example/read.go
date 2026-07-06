package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Read() {
	h := hetzner.NewEnvironment()

	for _, z := range h.Zones() {
		fmt.Printf(
			"Zone: %s (ttl=%d, records=%d, status=%s)\n",
			z.Name,
			z.TTL,
			z.RecordCount,
			z.Status,
		)

		for _, r := range h.Records(z) {
			values := join.CommaSpace(r.Values)

			if r.TTL != nil {
				fmt.Printf(
					"  %s %s %s (ttl=%d)\n",
					r.Type,
					r.Name,
					values,
					*r.TTL,
				)
			} else {
				fmt.Printf("  %s %s %s\n", r.Type, r.Name, values)
			}
		}
	}
}

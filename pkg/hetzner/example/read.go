package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Read() {
	h := hetzner.NewEnvironment()

	for _, z := range h.Zones() {
		console.Format(
			"Zone: %s (ttl=%d, records=%d, status=%s)\n",
			z.Name,
			z.TTL,
			z.RecordCount,
			z.Status,
		)

		for _, r := range h.Records(z) {
			values := join.CommaSpace(r.Values)

			if r.TTL != nil {
				console.Format(
					"  %s %s %s (ttl=%d)\n",
					r.Type,
					r.Name,
					values,
					*r.TTL,
				)
			} else {
				console.Format("  %s %s %s\n", r.Type, r.Name, values)
			}
		}
	}
}

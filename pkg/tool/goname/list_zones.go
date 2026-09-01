package goname

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/spf13/cobra"
)

func listZones(c *hetzner.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-zones",
		Short: "List all zones",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			for _, z := range c.Zones() {
				console.Format(
					"%s (ttl=%d, records=%d, status=%s)\n",
					z.Name,
					z.TTL,
					z.RecordCount,
					z.Status,
				)
			}
		},
	}
}

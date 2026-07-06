package gotechnitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func listZones(c *technitium.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-zones",
		Short: "List all authoritative zones",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			for _, z := range c.MustZones() {
				fmt.Printf(
					"%s (%s) internal=%v\n",
					z.Name,
					z.Type,
					z.Internal,
				)
			}
		},
	}
}

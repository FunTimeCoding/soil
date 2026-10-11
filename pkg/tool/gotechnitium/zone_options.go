package gotechnitium

import (
	"github.com/funtimecoding/soil/pkg/technitium"
	"github.com/spf13/cobra"
)

func zoneOptions(c *technitium.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "zone-options [zone]",
		Short: "Show a zone's dynamic update and zone transfer settings",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			printZoneOptions(c.MustZoneOptions(arguments[0]))
		},
	}
}

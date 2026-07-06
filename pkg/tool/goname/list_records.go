package goname

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/spf13/cobra"
)

func listRecords(c *hetzner.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-records [zone]",
		Short: "List records in a zone",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			z := zoneByName(c, arguments[0])

			for _, r := range c.Records(z) {
				fmt.Printf(
					"%s %s %s\n",
					r.Type,
					r.Name,
					join.CommaSpace(r.Values),
				)
			}
		},
	}
}

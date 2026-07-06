package goname

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/hetzner"
	"github.com/spf13/cobra"
)

func addRecord(c *hetzner.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-record [zone] [name] [type] [value]",
		Short: "Add a record - CNAME targets need the trailing dot",
		Args:  cobra.ExactArgs(4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			z := zoneByName(c, arguments[0])
			c.CreateRecord(
				z,
				arguments[1],
				arguments[2],
				arguments[3],
			)
			fmt.Printf(
				"added %s %s.%s -> %s\n",
				arguments[2],
				arguments[1],
				arguments[0],
				arguments[3],
			)
		},
	}
}

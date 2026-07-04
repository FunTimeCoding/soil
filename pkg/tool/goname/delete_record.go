package goname

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/hetzner"
	"github.com/spf13/cobra"
)

func deleteRecord(c *hetzner.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-record [zone] [name] [type]",
		Short: "Delete a record",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			z := zoneByName(c, arguments[0])
			c.DeleteRecord(z, arguments[1], arguments[2])
			fmt.Printf(
				"deleted %s %s.%s\n",
				arguments[2],
				arguments[1],
				arguments[0],
			)
		},
	}
}

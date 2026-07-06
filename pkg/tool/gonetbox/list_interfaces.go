package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listInterfaces(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-interfaces [device]",
		Short: "List interfaces for a NetBox device",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.ListInterfaces(arguments[0]))
		},
	}
}

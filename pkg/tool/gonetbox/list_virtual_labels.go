package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listVirtualLabels(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-virtual-labels [machine]",
		Short: "List labels on a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.ListVirtualLabels(arguments[0]))
		},
	}
}

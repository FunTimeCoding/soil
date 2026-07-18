package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func removeVirtualLabel(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-virtual-label [machine] [key]",
		Short: "Remove a label from a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			c.RemoveVirtualLabel(arguments[0], arguments[1])
			fmt.Println("label removed")
		},
	}
}

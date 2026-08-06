package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func renameVirtualMachine(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "rename-virtual-machine [name] [new-name]",
		Short: "Rename a NetBox virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.UpdateVirtualMachine(arguments[0], arguments[1], ""),
			)
		},
	}
}

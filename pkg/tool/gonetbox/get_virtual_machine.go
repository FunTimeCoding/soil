package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func getVirtualMachine(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-virtual-machine [name]",
		Short: "Get a NetBox virtual machine by name",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.GetVirtualMachine(arguments[0]))
		},
	}
}

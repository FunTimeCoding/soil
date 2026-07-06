package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createVirtualInterface(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-virtual-interface [vm] [name]",
		Short: "Create an interface on a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateVirtualInterface(arguments[0], arguments[1]))
		},
	}
}

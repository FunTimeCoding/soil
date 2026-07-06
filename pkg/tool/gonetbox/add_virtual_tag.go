package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func addVirtualTag(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "add-virtual-tag [vm] [tag]",
		Short: "Add a tag to a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.AddVirtualTag(arguments[0], arguments[1]))
		},
	}
}

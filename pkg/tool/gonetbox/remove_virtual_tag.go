package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func removeVirtualTag(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-virtual-tag [vm] [tag]",
		Short: "Remove a tag from a virtual machine",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.RemoveVirtualTag(arguments[0], arguments[1]))
		},
	}
}

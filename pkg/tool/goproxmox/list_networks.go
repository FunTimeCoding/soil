package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func listNetworks(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-networks [node]",
		Short: "List network interfaces on a node",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			console.Emit(c.Client().ListNetworks(a[0]))
		},
	}
}

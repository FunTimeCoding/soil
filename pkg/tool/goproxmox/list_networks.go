package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
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
			c.Terminal().EmitValue(
				c.Client().ListNetworks(
					context.Background(),
					a[0],
					&client.ListNetworksParams{Instance: c.Instance()},
				),
			)
		},
	}
}

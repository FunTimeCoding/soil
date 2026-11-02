package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func listContainers(c *command_context.Context) *cobra.Command {
	var node string
	result := &cobra.Command{
		Use:   "list-containers",
		Short: "List LXC containers",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var n *string

			if node != "" {
				n = &node
			}

			c.Terminal().EmitValue(
				c.Client().ListContainers(
					context.Background(),
					&client.ListContainersParams{
						Instance: c.Instance(),
						Node:     n,
					},
				),
			)
		},
	}
	result.Flags().StringVar(&node, "node", "", "filter by node name")

	return result
}

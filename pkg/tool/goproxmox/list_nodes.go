package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func listNodes(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-nodes",
		Short: "List Proxmox nodes",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().ListNodes(
					context.Background(),
					&client.ListNodesParams{Instance: c.Instance()},
				),
			)
		},
	}
}

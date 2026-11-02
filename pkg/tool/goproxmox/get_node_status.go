package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func getNodeStatus(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "get-node-status [name]",
		Short: "Get Proxmox node status",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			c.Terminal().EmitValue(
				c.Client().GetNodeStatus(
					context.Background(),
					a[0],
					&client.GetNodeStatusParams{Instance: c.Instance()},
				),
			)
		},
	}
}

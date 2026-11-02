package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func listInstances(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-instances",
		Short: "List configured Proxmox instances",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().ListInstances(context.Background()),
			)
		},
	}
}

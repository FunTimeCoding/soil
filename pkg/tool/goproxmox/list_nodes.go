package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
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
			console.Emit(c.Client().ListNodes())
		},
	}
}

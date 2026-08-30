package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func listStorages(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-storages [node]",
		Short: "List storage backends on a node",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			console.Emit(c.Client().ListStorages(a[0]))
		},
	}
}

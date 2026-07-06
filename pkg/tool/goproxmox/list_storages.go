package goproxmox

import (
	"fmt"
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
			fmt.Println(c.Client().ListStorages(a[0]))
		},
	}
}

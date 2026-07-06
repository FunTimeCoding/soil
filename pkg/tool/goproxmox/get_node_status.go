package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
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
			fmt.Println(c.Client().GetNodeStatus(a[0]))
		},
	}
}

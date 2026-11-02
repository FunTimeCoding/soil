package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func listSnapshots(c *command_context.Context) *cobra.Command {
	var node string
	result := &cobra.Command{
		Use:   "list-snapshots [identifier]",
		Short: "List snapshots for a virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			identifier, e := strconv.ParseInt(a[0], 10, 64)
			errors.PanicOnError(e)
			var n *string

			if node != "" {
				n = &node
			}

			c.Terminal().EmitValue(
				c.Client().ListMachineSnapshots(
					context.Background(),
					identifier,
					&client.ListMachineSnapshotsParams{
						Instance: c.Instance(),
						Node:     n,
					},
				),
			)
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name (speeds up lookup)")

	return result
}

package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func createContainerSnapshot(c *command_context.Context) *cobra.Command {
	var node string
	var name string
	result := &cobra.Command{
		Use:   "create-container-snapshot [identifier]",
		Short: "Create a container snapshot",
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
				c.Client().CreateContainerSnapshot(
					context.Background(),
					identifier,
					&client.CreateContainerSnapshotParams{
						Instance: c.Instance(),
						Node:     n,
					},
					client.CreateContainerSnapshotJSONRequestBody{Name: name},
				),
			)
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name")
	result.Flags().StringVar(&name, "name", "", "snapshot name")
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}

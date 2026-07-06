package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func listContainerSnapshots(c *command_context.Context) *cobra.Command {
	var node string
	result := &cobra.Command{
		Use:   "list-container-snapshots [identifier]",
		Short: "List snapshots for a container",
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

			fmt.Println(c.Client().ListContainerSnapshots(identifier, n))
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name")

	return result
}

package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteSnapshot(c *command_context.Context) *cobra.Command {
	var node string
	var name string
	result := &cobra.Command{
		Use:   "delete-snapshot [identifier]",
		Short: "Delete a snapshot from a virtual machine",
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

			fmt.Println(c.Client().DeleteMachineSnapshot(identifier, name, n))
		},
	}
	result.Flags().StringVar(
		&name,
		"name",
		"",
		"snapshot name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("name"))
	result.Flags().StringVar(
		&node,
		"node",
		"",
		"node name (speeds up lookup)",
	)

	return result
}

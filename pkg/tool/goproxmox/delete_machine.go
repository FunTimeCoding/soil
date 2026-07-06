package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
	"strconv"
)

func deleteMachine(c *command_context.Context) *cobra.Command {
	var node string
	var purge bool
	result := &cobra.Command{
		Use:   "delete-machine [identifier]",
		Short: "Delete a virtual machine",
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

			var p *bool

			if purge {
				p = &purge
			}

			fmt.Println(c.Client().DeleteMachine(identifier, n, p))
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name")
	result.Flags().BoolVar(
		&purge,
		"purge",
		false,
		"remove from cluster config",
	)

	return result
}

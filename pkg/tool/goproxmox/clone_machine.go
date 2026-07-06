package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func cloneMachine(c *command_context.Context) *cobra.Command {
	var node string
	var name string
	var full bool
	var storage string
	var snapshot string
	result := &cobra.Command{
		Use:   "clone-machine [identifier]",
		Short: "Clone a virtual machine",
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

			body := client.CloneMachineJSONRequestBody{
				Name: name,
			}

			if full {
				body.Full = &full
			}

			if storage != "" {
				body.Storage = &storage
			}

			if snapshot != "" {
				body.Snapshot = &snapshot
			}

			fmt.Println(c.Client().CloneMachine(identifier, n, body))
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name")
	result.Flags().StringVar(&name, "name", "", "name for the clone")
	result.Flags().BoolVar(&full, "full", false, "full clone")
	result.Flags().StringVar(&storage, "storage", "", "target storage")
	result.Flags().StringVar(&snapshot, "snapshot", "", "clone from snapshot")
	errors.PanicOnError(result.MarkFlagRequired("name"))

	return result
}

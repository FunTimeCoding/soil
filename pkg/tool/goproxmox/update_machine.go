package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
	"strconv"
)

func updateMachine(c *command_context.Context) *cobra.Command {
	var node string
	var name string
	var tags string
	var description string
	var deleteFields string
	result := &cobra.Command{
		Use:   "update-machine [identifier]",
		Short: "Update a virtual machine",
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

			var body client.UpdateMachineJSONRequestBody

			if name != "" {
				body.Name = &name
			}

			if tags != "" {
				body.Tags = &tags
			}

			if description != "" {
				body.Description = &description
			}

			if deleteFields != "" {
				body.Delete = &deleteFields
			}

			fmt.Println(c.Client().UpdateMachine(identifier, n, body))
		},
	}
	result.Flags().StringVar(&node, "node", "", "node name")
	result.Flags().StringVar(&name, "name", "", "new VM name")
	result.Flags().StringVar(&tags, "tags", "", "semicolon-separated tags")
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"VM description",
	)
	result.Flags().StringVar(&deleteFields, "delete", "", "fields to clear")

	return result
}

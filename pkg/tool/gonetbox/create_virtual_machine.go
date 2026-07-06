package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createVirtualMachine(c *client.Client) *cobra.Command {
	var cluster string
	result := &cobra.Command{
		Use:   "create-virtual-machine [name]",
		Short: "Create a NetBox virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateVirtualMachine(arguments[0], cluster))
		},
	}
	result.Flags().StringVar(
		&cluster,
		"cluster",
		"",
		"cluster name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("cluster"))

	return result
}

package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createVirtualMachine(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var cluster string
	result := &cobra.Command{
		Use:   "create-virtual-machine [name]",
		Short: "Create a NetBox virtual machine",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateVirtualMachine(
					context.Background(),
					arguments[0],
					cluster,
				),
			)
		},
	}
	result.Flags().StringVar(&cluster, "cluster", "", "cluster name (required)")
	errors.PanicOnError(result.MarkFlagRequired("cluster"))

	return result
}

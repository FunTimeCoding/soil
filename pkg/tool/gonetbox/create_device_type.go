package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createDeviceType(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var manufacturer string
	result := &cobra.Command{
		Use:   "create-device-type [model]",
		Short: "Create a NetBox device type",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateDeviceType(
					context.Background(),
					arguments[0],
					manufacturer,
				),
			)
		},
	}
	result.Flags().StringVar(
		&manufacturer,
		"manufacturer",
		"",
		"manufacturer name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("manufacturer"))

	return result
}

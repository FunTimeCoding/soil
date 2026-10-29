package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createAddressRange(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var status string
	var description string
	result := &cobra.Command{
		Use:   "create-address-range [start] [end]",
		Short: "Create an IP address range (addresses in CIDR notation)",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateAddressRange(
					context.Background(),
					arguments[0],
					arguments[1],
					status,
					description,
				),
			)
		},
	}
	result.Flags().StringVar(
		&status,
		"status",
		"",
		"range status (active, reserved, deprecated)",
	)
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"range description",
	)

	return result
}

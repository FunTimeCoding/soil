package gonetbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/spf13/cobra"
)

func createTunnel(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var encapsulation string
	var group string
	result := &cobra.Command{
		Use:   "create-tunnel [name]",
		Short: "Create a NetBox tunnel",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.EmitValue(
				c.CreateTunnel(
					context.Background(),
					arguments[0],
					encapsulation,
					group,
				),
			)
		},
	}
	result.Flags().StringVar(
		&encapsulation,
		"encapsulation",
		"",
		"encapsulation type (e.g. wireguard, gre, openvpn) (required)",
	)
	result.Flags().StringVar(
		&group,
		"group",
		"",
		"tunnel group name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("encapsulation"))
	errors.PanicOnError(result.MarkFlagRequired("group"))

	return result
}

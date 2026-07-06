package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createTunnel(c *client.Client) *cobra.Command {
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
			fmt.Println(c.CreateTunnel(arguments[0], encapsulation, group))
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

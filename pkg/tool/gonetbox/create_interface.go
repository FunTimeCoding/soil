package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createInterface(c *client.Client) *cobra.Command {
	var interfaceType string
	result := &cobra.Command{
		Use:   "create-interface [device] [name]",
		Short: "Create a network interface on a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateInterface(
					arguments[0],
					arguments[1],
					interfaceType,
				),
			)
		},
	}
	result.Flags().StringVar(
		&interfaceType,
		"type",
		"",
		"interface type (e.g. 1000base-t)",
	)
	errors.PanicOnError(result.MarkFlagRequired("type"))

	return result
}

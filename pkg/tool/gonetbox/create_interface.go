package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createInterface(c *client.Client) *cobra.Command {
	var interfaceType string
	var physicalAddress string
	result := &cobra.Command{
		Use:   "create-interface [device] [name]",
		Short: "Create a network interface on a device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(
				c.CreateInterface(
					arguments[0],
					arguments[1],
					interfaceType,
					physicalAddress,
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
	result.Flags().StringVar(
		&physicalAddress,
		"physical-address",
		"",
		"MAC address to create and assign to the interface",
	)
	errors.PanicOnError(result.MarkFlagRequired("type"))

	return result
}

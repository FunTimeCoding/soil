package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createVirtualAddress(c *client.Client) *cobra.Command {
	var interfaceName string
	result := &cobra.Command{
		Use:   "create-virtual-address [vm] [address]",
		Short: "Assign an IP address to a virtual machine interface",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateVirtualAddress(
					arguments[0],
					interfaceName,
					arguments[1],
				),
			)
		},
	}
	result.Flags().StringVar(
		&interfaceName,
		"interface",
		"",
		"interface name (required)",
	)
	errors.PanicOnError(result.MarkFlagRequired("interface"))

	return result
}

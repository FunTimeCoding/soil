package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createDeviceType(c *client.Client) *cobra.Command {
	var manufacturer string
	result := &cobra.Command{
		Use:   "create-device-type [model]",
		Short: "Create a NetBox device type",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.CreateDeviceType(arguments[0], manufacturer))
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

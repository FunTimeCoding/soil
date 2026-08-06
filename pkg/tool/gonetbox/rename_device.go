package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/spf13/cobra"
)

func renameDevice(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "rename-device [name] [new-name]",
		Short: "Rename a NetBox device",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.UpdateDevice(
					arguments[0],
					generated.UpdateDeviceRequest{Name: &arguments[1]},
				),
			)
		},
	}
}

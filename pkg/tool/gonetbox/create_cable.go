package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createCable(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-cable [device-a] [interface-a] [device-b] [interface-b]",
		Short: "Create a NetBox cable between two device interfaces",
		Args:  cobra.ExactArgs(4),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.CreateCable(
					arguments[0],
					arguments[1],
					arguments[2],
					arguments[3],
				),
			)
		},
	}
}

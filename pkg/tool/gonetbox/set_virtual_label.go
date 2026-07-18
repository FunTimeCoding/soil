package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func setVirtualLabel(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "set-virtual-label [machine] [key] [value]",
		Short: "Set a key-value label on a virtual machine",
		Args:  cobra.ExactArgs(3),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(
				c.SetVirtualLabel(
					arguments[0],
					arguments[1],
					arguments[2],
				),
			)
		},
	}
}

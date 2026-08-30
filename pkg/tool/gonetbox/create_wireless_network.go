package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func createWirelessNetwork(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "create-wireless-network [ssid]",
		Short: "Create a wireless network",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			console.Emit(c.CreateWirelessNetwork(arguments[0]))
		},
	}
}

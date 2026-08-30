package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listWirelessNetworks(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-wireless-networks",
		Short: "List all wireless networks",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListWirelessNetworks())
		},
	}
}

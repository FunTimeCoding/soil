package gonetbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listAddressRanges(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-address-ranges",
		Short: "List all IP address ranges",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListAddressRanges())
		},
	}
}

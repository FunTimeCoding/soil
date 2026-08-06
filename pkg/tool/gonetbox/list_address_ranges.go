package gonetbox

import (
	"fmt"
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
			fmt.Println(c.ListAddressRanges())
		},
	}
}

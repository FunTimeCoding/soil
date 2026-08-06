package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listLocations(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-locations",
		Short: "List all locations",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListLocations())
		},
	}
}

package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listCables(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-cables",
		Short: "List all NetBox cables",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListCables())
		},
	}
}

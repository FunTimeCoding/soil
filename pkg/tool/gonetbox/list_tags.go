package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listTags(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-tags",
		Short: "List all NetBox tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.ListTags())
		},
	}
}

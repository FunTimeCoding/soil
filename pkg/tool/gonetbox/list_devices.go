package gonetbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
)

func listDevices(c *client.Client) *cobra.Command {
	var query string
	result := &cobra.Command{
		Use:   "list-devices",
		Short: "List NetBox devices",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var q *string

			if query != "" {
				q = &query
			}

			fmt.Println(c.ListDevices(q))
		},
	}
	result.Flags().StringVar(
		&query,
		"query",
		"",
		"filter by name",
	)

	return result
}

package goatlassian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func searchUsers(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "search-users [query]",
		Short: "Search Jira users",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			fmt.Println(c.SearchUsers(arguments[0]))
		},
	}
}

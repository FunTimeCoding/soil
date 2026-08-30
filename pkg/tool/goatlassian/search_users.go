package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
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
			console.Emit(c.SearchUsers(arguments[0]))
		},
	}
}

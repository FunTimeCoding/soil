package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func listProjects(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-projects",
		Short: "List all visible Jira projects",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListProjects())
		},
	}
}

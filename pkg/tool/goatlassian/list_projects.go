package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func listProjects(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "list-projects",
		Short: "List all visible Jira projects",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.ListProjects())
		},
	}
}

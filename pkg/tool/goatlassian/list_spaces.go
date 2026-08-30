package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func listSpaces(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list-spaces",
		Short: "List all visible Confluence spaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.ListSpaces())
		},
	}
}

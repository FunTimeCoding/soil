package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func getLinkTypes(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "get-link-types",
		Short: "List available issue link types",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.GetLinkTypes())
		},
	}
}

package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func getSnippet(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "get-snippet [name]",
		Short: "Read a snippet file",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			c.Terminal().EmitValue(
				c.Client().GetSnippet(
					context.Background(),
					a[0],
					&client.GetSnippetParams{Instance: c.Instance()},
				),
			)
		},
	}
}

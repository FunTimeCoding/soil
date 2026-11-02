package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/spf13/cobra"
)

func listSnippets(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-snippets",
		Short: "List snippet files",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().ListSnippets(
					context.Background(),
					&client.ListSnippetsParams{Instance: c.Instance()},
				),
			)
		},
	}
}

package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
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
			fmt.Println(c.Client().ListSnippets())
		},
	}
}

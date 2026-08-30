package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
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
			console.Emit(c.Client().GetSnippet(a[0]))
		},
	}
}

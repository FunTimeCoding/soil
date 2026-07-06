package goproxmox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func deleteSnippet(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-snippet [name]",
		Short: "Delete a snippet file",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			fmt.Println(c.Client().DeleteSnippet(a[0]))
		},
	}
}

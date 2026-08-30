package goproxmox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/spf13/cobra"
)

func listStorageContent(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list-storage-content [node] [storage]",
		Short: "List content on a storage backend",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			a []string,
		) {
			console.Emit(c.Client().ListStorageContent(a[0], a[1]))
		},
	}
}

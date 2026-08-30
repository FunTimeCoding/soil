package gopnsense

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/spf13/cobra"
)

func interfaces(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "interfaces",
		Short: "List network interfaces",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Interfaces())
		},
	}
}

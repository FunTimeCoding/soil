package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func gear(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "gear",
		Short: "List owned gear",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Gear())
		},
	}
}

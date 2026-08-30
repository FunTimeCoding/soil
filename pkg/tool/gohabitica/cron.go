package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func cron(c *client.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "cron",
		Short: "Check and run daily rollover if needed",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			console.Emit(c.Cron())
		},
	}
}

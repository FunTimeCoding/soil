package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func cron(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "cron",
		Short: "Check and run daily rollover if needed",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.Cron())
		},
	}
}

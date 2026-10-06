package goraid

import (
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/spf13/cobra"
)

func reports(
	c *raid.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "reports",
		Short: "List generated reports",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.Reports())
		},
	}
}

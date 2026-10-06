package gohabitica

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func tags(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "tags",
		Short: "List tags",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			t.Emit(c.Tags())
		},
	}
}

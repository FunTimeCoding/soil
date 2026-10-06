package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func setPageStatus(
	c *client.Client,
	t *terminal.Terminal,
) *cobra.Command {
	return &cobra.Command{
		Use:   "set-page-status [identifier] [current|draft]",
		Short: "Publish or unpublish a Confluence page",
		Args:  cobra.ExactArgs(2),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			t.Emit(c.SetPageStatus(arguments[0], arguments[1]))
		},
	}
}

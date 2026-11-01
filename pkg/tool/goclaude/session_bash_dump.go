package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionBashDump(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "bash-dump",
		Short: "Dump all raw Bash commands across sessions",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			r, e := c.Client().GetSessionsBashDump(context.Background())

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			for _, command := range r.Commands {
				console.Line(command)
			}
		},
	}
}

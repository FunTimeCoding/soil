package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func usage(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "usage",
		Short: "Show current Claude usage",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			j, found, e := c.Client().GetUsage(context.Background())

			if e != nil {
				c.Terminal().EmitDone(e)

				return
			}

			if !found {
				c.Terminal().EmitDone(not_found.Format("no usage data yet"))

				return
			}

			c.Terminal().EmitValue(j, nil)
		},
	}
}

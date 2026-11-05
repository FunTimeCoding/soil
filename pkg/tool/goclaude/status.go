package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func status(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Report what is inconsistent inside goclauded",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(c.Client().GetStatus(context.Background()))
		},
	}
}

package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionPermissions(c *command_context.Context) *cobra.Command {
	var days int
	result := &cobra.Command{
		Use:   "permissions",
		Short: "Rank the last days' permission prompts by what blocked them",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Terminal().EmitValue(
				c.Client().GetPermissions(
					context.Background(),
					&client.GetPermissionsParams{Days: &days},
				),
			)
		},
	}
	result.Flags().IntVar(
		&days,
		"days",
		goclaudedConstant.PermissionWindowDays,
		"How many days back to count",
	)

	return result
}

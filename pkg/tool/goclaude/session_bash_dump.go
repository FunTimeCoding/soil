package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			response, e := c.Client().GetSessionsBashDumpWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)

			for _, command := range response.JSON200.Commands {
				fmt.Println(command)
			}
		},
	}
}

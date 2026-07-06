package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
	"os"
)

func wait(c *command_context.Context) *cobra.Command {
	var timeout int
	result := &cobra.Command{
		Use:   "wait",
		Short: "Long-poll goclauded for messages (asyncRewake hook)",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			name := os.Getenv(constant.NameEnvironment)
			response, e := c.Client().GetWaitWithResponse(
				context.Background(),
				&client.GetWaitParams{
					Callsign: name,
					Timeout:  &timeout,
				},
			)

			if e != nil {
				os.Exit(0)
			}

			if response.JSON200 == nil {
				os.Exit(0)
			}

			messages := response.JSON200.Messages

			if len(messages) == 0 {
				os.Exit(0)
			}

			for _, m := range messages {
				errors.Printf("%s: %s\n", m.From, m.Body)
			}

			os.Exit(2)
		},
	}
	result.Flags().IntVar(
		&timeout,
		"timeout",
		30,
		"poll timeout in seconds",
	)

	return result
}

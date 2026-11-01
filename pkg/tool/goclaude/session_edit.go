package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionEdit(c *command_context.Context) *cobra.Command {
	var name string
	var description string
	result := &cobra.Command{
		Use:   "edit <id-or-name>",
		Short: "Edit session metadata (name, description)",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				console.Format("session not found: %s\n", arguments[0])

				return
			}

			body := client.PostEditSessionJSONRequestBody{Session: identifier}

			if name != "" {
				body.Name = &name
			}

			if description != "" {
				body.Description = &description
			}

			if body.Name == nil && body.Description == nil {
				console.Line("provide --name or --description")

				return
			}

			if e := c.Client().PostEditSession(
				context.Background(),
				body,
			); e != nil {
				c.Terminal().Exitln(e)

				return
			}

			if name != "" {
				console.Format("renamed: %s → %s\n", identifier, name)
			}

			if description != "" {
				console.Format("described: %s\n", identifier)
			}
		},
	}
	result.Flags().StringVar(&name, "name", "", "Set display name")
	result.Flags().StringVar(&description, "description", "", "Set description")

	return result
}

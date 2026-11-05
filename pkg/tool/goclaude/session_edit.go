package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
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
			if name == "" && description == "" {
				c.Terminal().EmitDone(
					validation.New("provide --name or --description"),
				)

				return
			}

			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			body := client.PostEditSessionJSONRequestBody{Session: identifier}

			if name != "" {
				body.Name = &name
			}

			if description != "" {
				body.Description = &description
			}

			c.Terminal().EmitDone(
				c.Client().PostEditSession(context.Background(), body),
			)
		},
	}
	result.Flags().StringVar(&name, "name", "", "Set display name")
	result.Flags().StringVar(&description, "description", "", "Set description")

	return result
}

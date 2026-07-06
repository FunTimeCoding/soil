package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			body := client.PostEditSessionJSONRequestBody{
				Session: identifier,
			}

			if name != "" {
				body.Name = &name
			}

			if description != "" {
				body.Description = &description
			}

			if body.Name == nil && body.Description == nil {
				fmt.Println("provide --name or --description")

				return
			}

			response, e := c.Client().PostEditSessionWithResponse(
				context.Background(),
				body,
			)
			errors.PanicOnError(e)

			if response.JSON500 != nil {
				fmt.Println(response.JSON500.Error)

				return
			}

			if name != "" {
				fmt.Printf("renamed: %s → %s\n", identifier, name)
			}

			if description != "" {
				fmt.Printf("described: %s\n", identifier)
			}
		},
	}
	result.Flags().StringVar(
		&name,
		"name",
		"",
		"Set display name",
	)
	result.Flags().StringVar(
		&description,
		"description",
		"",
		"Set description",
	)

	return result
}

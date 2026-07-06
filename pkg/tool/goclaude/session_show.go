package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/spf13/cobra"
)

func sessionShow(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "show <id-or-name>",
		Short: "Show session detail with alias, description, completions, and summary",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf("session not found: %s\n", arguments[0])

				return
			}

			response, e := c.Client().GetSessionDetailWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)

			if response.JSON200 == nil {
				return
			}

			d := response.JSON200
			fmt.Printf("Identifier: %s\n", d.Identifier)

			if d.Name != nil {
				fmt.Printf("Name: %s\n", *d.Name)
			}

			if d.Callsign != nil {
				fmt.Printf("Callsign: %s\n", *d.Callsign)
			}

			if d.Alias != nil {
				fmt.Printf("Alias: %s\n", *d.Alias)
			}

			if d.Slug != nil {
				fmt.Printf("Slug: %s\n", *d.Slug)
			}

			if d.Created != nil {
				fmt.Printf("Created: %s\n", *d.Created)
			}

			if d.TurnCount != nil {
				fmt.Printf("Turns: %d\n", *d.TurnCount)
			}

			if d.Labels != nil && len(*d.Labels) > 0 {
				fmt.Println("\nLabels:")

				for _, l := range *d.Labels {
					fmt.Printf("  %s: %s\n", l.Key, l.Value)
				}
			}

			if d.Description != nil {
				fmt.Printf("\n%s\n", *d.Description)
			}

			if d.Completions != nil && len(*d.Completions) > 0 {
				fmt.Println("\nCompletions:")

				for _, o := range *d.Completions {
					fmt.Printf("  [%s] %s\n", o.Kind, o.Topic)

					if o.Summary != nil {
						fmt.Printf("    %s\n", *o.Summary)
					}
				}
			}

			if d.Summary != nil {
				fmt.Printf("\nSummary:\n%s\n", *d.Summary)
			}

			if d.Pulses != nil && len(*d.Pulses) > 0 {
				fmt.Println("\nPulses:")

				for _, p := range *d.Pulses {
					from := "→"

					if p.From != nil {
						from = fmt.Sprintf("%s →", *p.From)
					}

					fmt.Printf("  %s %s\n", from, p.Body)
				}
			}
		},
	}
}

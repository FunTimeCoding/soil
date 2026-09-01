package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
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
				console.Format("session not found: %s\n", arguments[0])

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
			console.Format("Identifier: %s\n", d.Identifier)

			if d.Name != nil {
				console.Format("Name: %s\n", *d.Name)
			}

			if d.Callsign != nil {
				console.Format("Callsign: %s\n", *d.Callsign)
			}

			if d.Alias != nil {
				console.Format("Alias: %s\n", *d.Alias)
			}

			if d.Slug != nil {
				console.Format("Slug: %s\n", *d.Slug)
			}

			if d.Created != nil {
				console.Format("Created: %s\n", *d.Created)
			}

			if d.TurnCount != nil {
				console.Format("Turns: %d\n", *d.TurnCount)
			}

			if d.Cost != nil {
				console.Format("Cost: $%.2f\n", *d.Cost)
			}

			if d.Usage != nil && len(*d.Usage) > 0 {
				console.Line("\nUsage:")

				for _, u := range *d.Usage {
					marker := ""

					if !pricing.KnownModel(u.Model) {
						marker = " (unknown model, sonnet rates)"
					}

					console.Format(
						"  %-8s %d calls, %s input, %s output, %s cache-write, %s cache-read, $%.2f%s\n",
						u.Model,
						u.Calls,
						pricing.FormatTokens(u.Input),
						pricing.FormatTokens(u.Output),
						pricing.FormatTokens(u.CacheCreation),
						pricing.FormatTokens(u.CacheRead),
						u.Cost,
						marker,
					)
				}
			}

			if d.Labels != nil && len(*d.Labels) > 0 {
				console.Line("\nLabels:")

				for _, l := range *d.Labels {
					console.Format("  %s: %s\n", l.Key, l.Value)
				}
			}

			if d.Description != nil {
				console.Format("\n%s\n", *d.Description)
			}

			if d.Completions != nil && len(*d.Completions) > 0 {
				console.Line("\nCompletions:")

				for _, o := range *d.Completions {
					console.Format("  [%s] %s\n", o.Kind, o.Topic)

					if o.Summary != nil {
						console.Format("    %s\n", *o.Summary)
					}
				}
			}

			if d.Summary != nil {
				console.Format("\nSummary:\n%s\n", *d.Summary)
			}

			if d.Pulses != nil && len(*d.Pulses) > 0 {
				console.Line("\nPulses:")

				for _, p := range *d.Pulses {
					from := "→"

					if p.From != nil {
						from = fmt.Sprintf("%s →", *p.From)
					}

					console.Format("  %s %s\n", from, p.Body)
				}
			}
		},
	}
}

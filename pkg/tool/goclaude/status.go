package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
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
			response, e := c.Client().GetStatusWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)

			if response.JSON200 == nil {
				console.Format(
					"unexpected response: %s\n",
					response.HTTPResponse.Status,
				)

				return
			}

			if len(response.JSON200.Findings) == 0 {
				console.Line("nothing inconsistent")

				return
			}

			for _, i := range response.JSON200.Findings {
				if i.Subject == nil {
					console.Format("%-24s %s\n", i.Kind, i.Detail)

					continue
				}

				console.Format("%-24s %s: %s\n", i.Kind, *i.Subject, i.Detail)
			}
		},
	}
}

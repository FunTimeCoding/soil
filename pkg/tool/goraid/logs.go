package goraid

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/spf13/cobra"
	"time"
)

func logs(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	var offset int
	var limit int
	var start string
	var end string
	result := &cobra.Command{
		Use:   "logs",
		Short: "List available combat logs",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			var o *int
			var l *int
			var s *time.Time
			var e *time.Time

			if offset > 0 {
				o = &offset
			}

			if limit > 0 {
				l = &limit
			}

			if start != "" {
				parsed, f := time.Parse(time.RFC3339, start)

				if f == nil {
					s = &parsed
				}
			}

			if end != "" {
				parsed, f := time.Parse(time.RFC3339, end)

				if f == nil {
					e = &parsed
				}
			}

			t.EmitValue(
				c.GetLogs(
					context.Background(),
					&client.GetLogsParams{
						Offset: o,
						Limit:  l,
						Start:  s,
						End:    e,
					},
				),
			)
		},
	}
	result.Flags().IntVar(&offset, "offset", 0, "offset into log list")
	result.Flags().IntVar(
		&limit,
		"limit",
		0,
		"maximum number of logs to return",
	)
	result.Flags().StringVar(
		&start,
		"start",
		"",
		"filter logs after this time (RFC3339)",
	)
	result.Flags().StringVar(
		&end,
		"end",
		"",
		"filter logs before this time (RFC3339)",
	)

	return result
}

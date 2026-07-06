package goraid

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/spf13/cobra"
	"time"
)

func logs(c *raid.Client) *cobra.Command {
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
				t, f := time.Parse(time.RFC3339, start)

				if f == nil {
					s = &t
				}
			}

			if end != "" {
				t, f := time.Parse(time.RFC3339, end)

				if f == nil {
					e = &t
				}
			}

			fmt.Println(c.Logs(o, l, s, e))
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

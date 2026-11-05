package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
)

func sessionPeek(c *command_context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "peek <id-or-name>",
		Short: "Scan a session for naming",
		Args:  cobra.ExactArgs(1),
		Run: func(
			_ *cobra.Command,
			arguments []string,
		) {
			identifier := resolveSession(c, arguments[0])

			if identifier == "" {
				return
			}

			p, e := c.Client().GetSessionPeek(context.Background(), identifier)

			if e != nil {
				c.Terminal().Exitln(e)

				return
			}

			t := c.Terminal()
			lines := p.LineCount
			total := len(p.Entries)
			t.Line(
				fmt.Sprintf(
					"%d lines, %d user messages",
					lines,
					p.UserMessageCount,
				),
			)
			printToolSummary(t, p.TotalToolCalls, p.ToolCounts)
			t.Line()

			if total == 0 {
				return
			}

			n := peekBlockSize(lines)
			midLimit := (peekContextDepth(lines) - 1) * 100

			if total*2 <= constant.PeekOutputBudget {
				for i, entry := range p.Entries {
					limit := midLimit

					if i == total-1 {
						limit = constant.PeekContextLimit
					}

					printPeekEntry(
						t,
						entry.UserText,
						entry.AssistantContext,
						limit,
					)
				}

				return
			}

			head := n

			if head > total {
				head = total
			}

			tail := n

			if head+tail > total {
				tail = total - head
			}

			middle := total - head - tail
			budget := constant.PeekOutputBudget - (head+tail)*2
			samples := 0

			if budget > 0 && middle > 0 {
				samples = budget / 2

				if samples > middle {
					samples = middle
				}
			}

			for _, entry := range p.Entries[:head] {
				printPeekEntry(
					t,
					entry.UserText,
					entry.AssistantContext,
					midLimit,
				)
			}

			if samples > 0 && middle > 0 {
				step := middle / samples
				t.Line(fmt.Sprintf("--- (%d messages) ---", middle))

				for i := 0; i < samples; i++ {
					idx := head + i*step
					entry := p.Entries[idx]
					printPeekEntry(
						t,
						entry.UserText,
						entry.AssistantContext,
						midLimit,
					)
				}
			} else if middle > 0 {
				t.Line(fmt.Sprintf("--- (%d messages) ---", middle))
			}

			if tail > 0 {
				t.Line("---")

				for j, entry := range p.Entries[total-tail:] {
					limit := midLimit

					if j == tail-1 {
						limit = constant.PeekContextLimit
					}

					printPeekEntry(
						t,
						entry.UserText,
						entry.AssistantContext,
						limit,
					)
				}
			}
		},
	}
}

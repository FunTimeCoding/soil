package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
			identifier := resolveSession(c.Client(), arguments[0])

			if identifier == "" {
				fmt.Printf(
					"session not found: %s\n",
					arguments[0],
				)

				return
			}

			response, e := c.Client().GetSessionPeekWithResponse(
				context.Background(),
				identifier,
			)
			errors.PanicOnError(e)
			p := response.JSON200
			lines := p.LineCount
			total := len(p.Entries)
			fmt.Printf(
				"%d lines, %d user messages\n",
				lines,
				p.UserMessageCount,
			)
			printToolSummary(p.TotalToolCalls, p.ToolCounts)
			fmt.Println()

			if total == 0 {
				return
			}

			n := peekBlockSize(lines)
			m := peekContextDepth(lines)

			if total*2 <= constant.PeekOutputBudget {
				for _, entry := range p.Entries {
					printPeekEntry(entry.UserText, entry.AssistantContext, m)
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
				printPeekEntry(entry.UserText, entry.AssistantContext, m)
			}

			if samples > 0 && middle > 0 {
				step := middle / samples
				fmt.Printf("--- (%d messages) ---\n", middle)

				for i := 0; i < samples; i++ {
					idx := head + i*step
					entry := p.Entries[idx]
					printPeekEntry(entry.UserText, entry.AssistantContext, m)
				}
			} else if middle > 0 {
				fmt.Printf("--- (%d messages) ---\n", middle)
			}

			if tail > 0 {
				fmt.Printf("---\n")

				for _, entry := range p.Entries[total-tail:] {
					printPeekEntry(entry.UserText, entry.AssistantContext, m)
				}
			}
		},
	}
}

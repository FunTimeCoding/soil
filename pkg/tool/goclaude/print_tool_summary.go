package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"strings"
)

func printToolSummary(
	total int,
	counts []client.ToolCount,
) {
	if total == 0 {
		return
	}

	top := 3

	if len(counts) < top {
		top = len(counts)
	}

	var parts []string

	for _, tc := range counts[:top] {
		parts = append(
			parts,
			fmt.Sprintf("%d %s", tc.Count, tc.Name),
		)
	}

	fmt.Printf(
		"%d tool calls (%s)\n",
		total,
		strings.Join(parts, ", "),
	)
}

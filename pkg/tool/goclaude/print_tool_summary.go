package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func printToolSummary(
	total int,
	counts []client.ToolCount,
) {
	if total == 0 {
		return
	}

	top := min(len(counts), 3)
	var parts []string

	for _, tc := range counts[:top] {
		parts = append(parts, fmt.Sprintf("%d %s", tc.Count, tc.Name))
	}

	fmt.Printf("%d tool calls (%s)\n", total, join.CommaSpace(parts))
}

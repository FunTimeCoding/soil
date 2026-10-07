package claude

import (
	"github.com/funtimecoding/soil/pkg/generative/types/session_tool_count"
	"sort"
	"strings"
)

func (c *Client) SessionsByTool(toolFilter string) []*session_tool_count.Count {
	var result []*session_tool_count.Count

	for _, s := range c.Sessions() {
		count := 0

		for _, call := range c.ToolCalls(s.Identifier) {
			if strings.Contains(call.Name, toolFilter) {
				count++
			}
		}

		if count == 0 {
			continue
		}

		result = append(result, session_tool_count.New(s, count))
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Count > result[j].Count
		},
	)

	return result
}

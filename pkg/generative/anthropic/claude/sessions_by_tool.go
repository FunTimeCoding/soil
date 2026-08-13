package claude

import (
	"sort"
	"strings"
)

func (c *Client) SessionsByTool(toolFilter string) []*SessionToolCount {
	var result []*SessionToolCount

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

		result = append(result, &SessionToolCount{Session: s, Count: count})
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Count > result[j].Count
		},
	)

	return result
}

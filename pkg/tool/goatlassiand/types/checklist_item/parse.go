package checklist_item

import "strings"

func Parse(value string) []*Item {
	var result []*Item

	for _, line := range strings.Split(value, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "+ ") {
			result = append(result, New(len(result)+1, line[2:], true))
		} else if strings.HasPrefix(line, "- ") {
			result = append(result, New(len(result)+1, line[2:], false))
		}
	}

	return result
}

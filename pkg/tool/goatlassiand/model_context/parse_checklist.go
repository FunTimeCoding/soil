package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"strings"
)

func ParseChecklist(value string) []*checklist_item.Item {
	var result []*checklist_item.Item

	for _, line := range strings.Split(value, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "+ ") {
			result = append(
				result,
				checklist_item.New(
					len(result)+1,
					line[2:],
					true,
				),
			)
		} else if strings.HasPrefix(line, "- ") {
			result = append(
				result,
				checklist_item.New(
					len(result)+1,
					line[2:],
					false,
				),
			)
		}
	}

	return result
}

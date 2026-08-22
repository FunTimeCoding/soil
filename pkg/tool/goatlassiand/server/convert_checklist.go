package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

func convertChecklist(items []*checklist_item.Item) []*server.ChecklistItem {
	result := make([]*server.ChecklistItem, 0, len(items))

	for _, i := range items {
		result = append(
			result,
			&server.ChecklistItem{
				Index:   i.Index,
				Text:    i.Text,
				Checked: i.Checked,
			},
		)
	}

	return result
}

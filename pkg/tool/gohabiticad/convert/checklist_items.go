package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/checklist_item"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func ChecklistItems(v []*checklist_item.Item) *[]*server.ChecklistItem {
	result := make([]*server.ChecklistItem, 0, len(v))

	for _, i := range v {
		result = append(result, ChecklistItem(i))
	}

	return &result
}

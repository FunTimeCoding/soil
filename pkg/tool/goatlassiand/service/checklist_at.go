package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

func checklistAt(
	items []*checklist_item.Item,
	index int,
) error {
	if index < 1 || index > len(items) {
		return not_found.Format(
			"index %d out of range (1-%d)",
			index,
			len(items),
		)
	}

	return nil
}

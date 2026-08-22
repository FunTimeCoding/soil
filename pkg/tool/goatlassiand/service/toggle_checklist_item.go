package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"

func (s *Service) ToggleChecklistItem(
	key string,
	index int,
) ([]*checklist_item.Item, error) {
	items, e := s.ReadChecklist(key)

	if e != nil {
		return nil, e
	}

	if f := checklistAt(items, index); f != nil {
		return nil, f
	}

	items[index-1].Checked = !items[index-1].Checked

	if f := s.WriteChecklist(key, items); f != nil {
		return nil, f
	}

	return items, nil
}
